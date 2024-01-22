function uploadCart() {

    var gameID = Number(document.getElementById('gameID').innerHTML);

    var settings = {
        "url": "http://localhost:8080/cart/upload",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "gameID": gameID,
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            window.location.href = 'cart';
        } else if (response["status"] == "SQL Access Error") {
            alert("Already in the shopping cart or inventory!");
        } else {
            uploadTempCart()
        }

    });
}

function uploadTempCart() {

    var tempCartJSON = localStorage.getItem("tempCart");
    var tempCart = JSON.parse(tempCartJSON);

    var tempDateJSON = localStorage.getItem("tempDate");
    var tempDate = JSON.parse(tempDateJSON);

    if (tempCart == null) {
        tempCart = [];
    }

    if (tempDate == null) {
        tempDate = {};
    }

    if (tempCart.includes(Number(document.getElementById("gameID").innerText)) == true) {
        tempCart = tempCart.filter(item => item !== Number(document.getElementById("gameID").innerText));
        alert("Already in the shopping cart or inventory!");
    }

    tempCart.push(Number(document.getElementById("gameID").innerText))

    var currentDateTime = new Date();

    var currentYear = currentDateTime.getFullYear(); // 获取当前年份
    var currentMonth = currentDateTime.getMonth() + 1; // 获取当前月份（注意要加1，因为月份从0开始）
    var currentDay = currentDateTime.getDate(); // 获取当前日期
    var currentHour = currentDateTime.getHours(); // 获取当前小时数
    var currentMinute = currentDateTime.getMinutes(); // 获取当前分钟数
    var currentSecond = currentDateTime.getSeconds(); // 获取当前秒数

    var now = currentYear + '-' + currentMonth + '-' + currentDay + ' ' + currentHour + ':' + currentMinute + ':' + currentSecond;
    tempDate[document.getElementById("gameID").innerText] = now;

    var tempCartJSON = JSON.stringify(tempCart);
    localStorage.setItem("tempCart", tempCartJSON);

    var tempDateJSON = JSON.stringify(tempDate);
    localStorage.setItem("tempDate", tempDateJSON);

    console.log(tempCart)
    console.log(tempDate)

    window.location.href = 'cart';
}

function browserCart() {

    var settings = {
        "url": "http://localhost:8080/cart/browser",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            console.log(response)

            var totalPrice = 0

            for (var i = 0; i < response["cartList"].length; i++) {

                totalPrice += Number(response["gameList"][i]["gamePrice"])

                //创建DIV元素
                var div = $('<div>').addClass('item');

                //创建UL元素
                var ul = $('<ul>');

                //创建第一个LI元素
                var li1 = $('<li>');
                var img = $('<img>').attr('src', response["gameList"][i]["gameImg"]).addClass('templatemo-item');
                li1.append(img);
                ul.append(li1);

                //创建第二个LI元素
                var li2 = $('<li>');
                var h4_1 = $('<h4>').text(response["gameList"][i]["gameName"]);
                var span_1 = $('<span>').text(response["gameList"][i]["gameUploader"]);
                li2.append(h4_1).append(span_1);
                ul.append(li2);

                //创建第三个LI元素
                var li3 = $('<li>');
                var h4_2 = $('<h4>').text('Date Added');
                var span_2 = $('<span>').text(response["cartList"][i]["cartDateAdded"]);
                li3.append(h4_2).append(span_2);
                ul.append(li3);

                //创建第四个LI元素
                var li4 = $('<li>');
                var h4_3 = $('<h4>').text('Price');
                var span_3 = $('<span>').text(response["gameList"][i]["gamePrice"]);
                li4.append(h4_3).append(span_3);
                ul.append(li4);

                //创建第五个LI元素
                var li5 = $('<li>');
                var h4_4 = $('<h4>').text('Current Sales');
                var span_4 = $('<span>').text('0');
                li5.append(h4_4).append(span_4);
                ul.append(li5);

                //创建第六个LI元素
                var li6 = $('<li>');
                // var div_1 = $('<div>').addClass('border-button');
                var div_1 = $('<div>', {
                    class: 'border-button',
                    id: response["gameList"][i]["gameID"],
                    onclick: 'removeCart(this.id)'
                });
                var a = $('<a>').attr('href', '#').text('Remove it');
                div_1.append(a);
                li6.append(div_1);
                ul.append(li6);

                div.append(ul);

                //将创建的DIV元素添加到DOM中

                $('#cartList').append(div);
            }

            document.getElementById('totalPrice').innerHTML = totalPrice
            document.getElementById('gameNumber').innerHTML = response["cartList"].length

        } else {
            browserTempCart()
        }

    });
}

function browserTempCart() {

    var tempCartJSON = localStorage.getItem("tempCart");
    var tempCart = JSON.parse(tempCartJSON);

    var tempDateJSON = localStorage.getItem("tempDate");
    var tempDate = JSON.parse(tempDateJSON);


    if (tempCart == null) {
        tempCart = [];
    }

    console.log(tempCart)

    var settings = {
        "url": "http://localhost:8080/cart/browserTemp",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            gameID: tempCart
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            console.log(response)

            var totalPrice = 0

            for (var i = 0; i < response["gameList"].length; i++) {

                totalPrice += Number(response["gameList"][i]["gamePrice"])

                //创建DIV元素
                var div = $('<div>').addClass('item');

                //创建UL元素
                var ul = $('<ul>');

                //创建第一个LI元素
                var li1 = $('<li>');
                var img = $('<img>').attr('src', response["gameList"][i]["gameImg"]).addClass('templatemo-item');
                li1.append(img);
                ul.append(li1);

                //创建第二个LI元素
                var li2 = $('<li>');
                var h4_1 = $('<h4>').text(response["gameList"][i]["gameName"]);
                var span_1 = $('<span>').text(response["gameList"][i]["gameUploader"]);
                li2.append(h4_1).append(span_1);
                ul.append(li2);

                //创建第三个LI元素
                var li3 = $('<li>');
                var h4_2 = $('<h4>').text('Date Added');
                var span_2 = $('<span>').text(tempDate[response["gameList"][i]["gameID"]]);
                li3.append(h4_2).append(span_2);
                ul.append(li3);

                //创建第四个LI元素
                var li4 = $('<li>');
                var h4_3 = $('<h4>').text('Price');
                var span_3 = $('<span>').text(response["gameList"][i]["gamePrice"]);
                li4.append(h4_3).append(span_3);
                ul.append(li4);

                //创建第五个LI元素
                var li5 = $('<li>');
                var h4_4 = $('<h4>').text('Current Sales');
                var span_4 = $('<span>').text('0');
                li5.append(h4_4).append(span_4);
                ul.append(li5);

                //创建第六个LI元素
                var li6 = $('<li>');
                // var div_1 = $('<div>').addClass('border-button');
                var div_1 = $('<div>', {
                    class: 'border-button',
                    id: response["gameList"][i]["gameID"],
                    onclick: 'removeCart(this.id)'
                });
                var a = $('<a>').attr('href', '#').text('Remove it');
                div_1.append(a);
                li6.append(div_1);
                ul.append(li6);

                div.append(ul);

                //将创建的DIV元素添加到DOM中

                $('#cartList').append(div);
            }

            document.getElementById('totalPrice').innerHTML = totalPrice
            document.getElementById('gameNumber').innerHTML = response["gameList"].length

        } else {
            alert("Local Storage Loading Fail!")
        }

    });

}

function removeCart(buttonId) {

    const gameID = Number(buttonId)
    console.log(gameID)

    var settings = {
        "url": "http://localhost:8080/cart/remove",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "gameID": gameID,
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            alert("Success!");
            location.reload();
        } else {

            var tempCartJSON = localStorage.getItem("tempCart");
            var tempCart = JSON.parse(tempCartJSON);

            if (tempCart.includes(Number(buttonId)) == true) {
                tempCart = tempCart.filter(item => item !== Number(buttonId));
                // alert("Already in the shopping cart or inventory!");
            }

            var tempCartJSON = JSON.stringify(tempCart);
            localStorage.setItem("tempCart", tempCartJSON);

            location.reload();
        }

    });
}

function checkCart() {

    var settings = {
        "url": "http://localhost:8080/cart/browser",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            console.log(response)

            var totalPrice = 0

            for (var i = 0; i < response["gameList"].length; i++) {

                console.log(response["gameList"].length)
                console.log(i)

                var li = $("<li></li>");
                li.text(response["gameList"][i]["gameName"]);

                var span = $("<span></span>");
                var icon = $("<i></i>");
                icon.addClass("fa fa-usd");
                span.append(icon);
                span.append(' ' + response["gameList"][i]["gamePrice"]);

                li.append(span);

                // 将新的 <li> 元素添加到现有的列表中
                $('#checkList').append(li);

                totalPrice += Number(response["gameList"][i]["gamePrice"])
            }

            document.getElementById('totalPrice').innerHTML = totalPrice

            window.location.href = 'payment';

        } else {
            alert("Please Login in first!");
            window.location.href = 'login';
        }

    });
}

function checkPayment() {

    var settings = {
        "url": "http://localhost:8080/cart/check",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
        }),
    };

    $.ajax(settings).done(function (response) {
        if (response["status"] == "Accepted") {
            window.location.href = 'result';
        } else {
            alert("Fail!");
        }

    });
}

function syncCart() {

    var tempCartJSON = localStorage.getItem("tempCart");
    var tempCart = JSON.parse(tempCartJSON);

    var tempDateJSON = localStorage.getItem("tempDate");
    var tempDate = JSON.parse(tempDateJSON);

    var tempDateList = []

    if (tempCart == null) {
        window.location.href = 'index';
        return
    }

    for (var i = 0; i < tempCart.length; i++) {
        tempDateList.push(tempDate[tempCart[i]])
    }

    // console.log(tempCart)
    // console.log(tempDateList)


    var settings = {
        "url": "http://localhost:8080/cart/sync",
        "method": "POST",
        "timeout": 0,
        "headers": {
            "Content-Type": "application/json"
        },
        "data": JSON.stringify({
            "gameID": tempCart,
            "cartDateAdded": tempDateList
        }),
    };

    $.ajax(settings).done(function (response) {

        console.log(response)

        if (response["status"] == "Accepted") {

            window.location.href = 'index';

        } else if (response["status"] == "CartSync") {

            alert("The cart synchronization is complete!")

            window.location.href = 'cart';

        } else {
            alert("Login in Fail!")
        }

    });
}