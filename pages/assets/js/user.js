function djb2Code(str) {
  var hash = 5381;
  for (var i = 0; i < str.length; i++) {
    hash = ((hash << 5) + hash) + str.charCodeAt(i);
  }
  return hash;
}

// set cookie
function setCookie(cname, cvalue, exdays) {
  var d = new Date();
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  var expires = "expires=" + d.toUTCString();
  document.cookie = cname + "=" + cvalue + "; " + expires;
}

function regist() {
  var Account = document.forms["registTable"]["registAccount"].value;
  var Password = document.forms["registTable"]["registPassword"].value;
  var Name = document.forms["registTable"]["registName"].value;

  if (Account == null || Account == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the Account";
    return false;
  }

  if (Password == null || Password == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the Password";
    return false;
  }

  if (Name == null || Name == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the User Name";
    return false;
  }


  var settings = {
    "url": "http://localhost:8080/user/regist",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json"
    },
    "data": JSON.stringify({
      "userAccount": Account,
      "userPassword": Password,
      "userName": Name,
    }),
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Account already exists") {
      document.getElementById("loginError").innerHTML =
        "Account already exists";
    } else {
      alert("Sign up Success!");
      window.location.href = 'login';
    }

    console.log(response);
  });

}

function login() {
  var Account = document.forms["loginTable"]["loginAccount"].value;
  var Password = document.forms["loginTable"]["loginPassword"].value;

  if (Account == null || Account == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the Account";
    return false;
  }

  if (Password == null || Password == "") {
    document.getElementById("loginError").innerHTML =
      "Please enter the Password";
    return false;
  }

  var settings = {
    "url": "http://localhost:8080/user/login",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json"
    },
    "data": JSON.stringify({
      "userAccount": Account,
      "userPassword": Password,
    }),
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Accepted") {

      alert("Login in Success!");

      syncCart()

    } else {
      document.getElementById("loginError").innerHTML =
        "Wrong Username or Password";
    }

    console.log(response);
  });
}

function logout() {
  var settings = {
    "url": "http://localhost:8080/user/logout",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {
    if (response["status"] == "Accepted") {

      localStorage.clear();
      document.cookie = 'sessionID=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';

      alert("Logout success!")
    } else {
      alert("Logout fali!")
    }
    console.log(response);
  });
}

function profileload() {

  var settings = {
    "url": "http://localhost:8080/user/profile",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {

    console.log(response)

    if (response["status"] == "Accepted") {
      document.getElementById("userName").innerHTML =
        response["userName"];
      document.getElementById("userGameCount").innerHTML =
        response["userGameCount"];
      if (response["userPortrait"] != '')
        document.getElementById("userPortrait").src = response["userPortrait"];

    } else {
      alert("Please log in first!")
      window.location.href = 'login';
    }
  });
}

function getInventory() {

  var settings = {
    "url": "http://localhost:8080/user/inventory",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {

    if (response["status"] == "Accepted") {

      console.log(response)

      for (var j = 0; j < response["gameList"].length; j++) {

        // 创建一个 div 元素，添加类名 "item"
        var $item = $("<div>").addClass("item");

        // 创建一个 ul 元素
        var $ul = $("<ul>");

        // 添加第一个 li 元素到 ul 中
        $("<li>").append($("<img>").attr("src", response["itemList"][j]["gameImg"]).attr("alt", "").addClass("templatemo-item")).appendTo($ul);

        // 添加第二个 li 元素到 ul 中
        $("<li>").append($("<h4>").text(response["itemList"][j]["gameName"])).append($("<span>").text(response["itemList"][j]["gameUploader"])).appendTo($ul);

        // 添加第三个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Date Added")).append($("<span>").text(response["gameList"][j]["InventoryDateAdded"])).appendTo($ul);

        // 添加第四个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Hours Played")).append($("<span>").text(response["gameList"][j]["InventoryHoursPlayed"])).appendTo($ul);

        // 添加第五个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Currently")).append($("<span>").text("Not downloaded")).appendTo($ul);

        // 添加第六个 li 元素到 ul 中
        $("<li>").append($("<div>").addClass("main-border-button").append($("<a>").attr("href", "#").text("Download"))).appendTo($ul);

        // 将 ul 元素添加到 div 中
        $item.append($ul);

        // 将生成的 div 元素添加到页面中
        $("#row").append($item);

      }

    } else {
      alert("Fail!")
    }

  });
}

function getHistory() {

  var settings = {
    "url": "http://localhost:8080/user/inventory",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "application/json",
    },
  };

  $.ajax(settings).done(function (response) {

    if (response["status"] == "Accepted") {

      console.log(response)

      for (var j = 0; j < response["gameList"].length; j++) {

        // 创建一个 div 元素，添加类名 "item"
        var $item = $("<div>").addClass("item");

        // 创建一个 ul 元素
        var $ul = $("<ul>");

        // 添加第一个 li 元素到 ul 中
        $("<li>").append($("<img>").attr("src", response["itemList"][j]["gameImg"]).attr("alt", "").addClass("templatemo-item")).appendTo($ul);

        // 添加第二个 li 元素到 ul 中
        $("<li>").append($("<h4>").text(response["itemList"][j]["gameName"])).append($("<span>").text(response["itemList"][j]["gameUploader"])).appendTo($ul);

        // 添加第三个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Date Added")).append($("<span>").text(response["gameList"][j]["InventoryDateAdded"])).appendTo($ul);

        // 添加第五个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Currently")).append($("<span>").text("In inventory")).appendTo($ul);

        // 添加第四个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Price")).append($("<span>").text(response["itemList"][j]["gamePrice"])).appendTo($ul);

        console.log(window.location.hash)

        // 添加第六个 li 元素到 ul 中
        $("<li>").append($("<h4>").text("Transaction number")).append($("<span>").text(djb2Code(response["gameList"][j]["userID"] + response["gameList"][j]["gameID"] + response["itemList"][j]["gameName"]))).appendTo($ul);

        // 将 ul 元素添加到 div 中
        $item.append($ul);

        // 将生成的 div 元素添加到页面中
        $("#row").append($item);

      }

    } else {
      alert("Fail!")
    }

  });
}

