$("#login").click(function(){
    let BNJ_USERNAME = $("#BNJ_USERNAME").val();
    let BNJ_PWD = $("#BNJ_PWD").val();
    if(!BNJ_USERNAME) {
        alert("請輸入帳號");
        return
    }
    if(!BNJ_PWD) {
        alert("請輸入密碼");
        return
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: '/login',
        data: { Username: BNJ_USERNAME, Password: BNJ_PWD },
        success: function (result) {
            if(!result.error) {
                alert(result.msg);
                window.location.href = "";
            } else {
                alert(result.msg);
            }
        }
    });
})