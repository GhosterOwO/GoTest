$("#logout").click(function(){
    $.ajax({
        type: "POST",
        dataType: "json",
        url: '/logout',
        success: function (result) {
            if(!result.error) {
                alert(result.msg);
                window.location.href = "";
                return
            }
            alert(result.msg);
        }
    });
})

function getImg() {
    $.ajax({
        type: "GET",
        dataType: "json",
        url: '/getimg',
        success: function (result) {
            if(!result.error) {
                $("#img *").remove();
                const data = result.data;
                let html = "";
                for(let i=0;i<data.length;i++) {
                    html += "<div><span>" + data[i] + "</span><img src='public/img/" + data[i] + "'><button class='deleteImg'>刪除</button></div><br/>"
                }
                $("#img").append(html);
                deleteImgBtn();
                return
            }
        }
    });
}

function deleteImgBtn(){
    $(".deleteImg").off().click(function(){
        const ImgName = $(this).prevAll("span").text();
        const div = $(this).parent();
        const br = $(this).parent().next();
        $.ajax({
            type: "POST",
            dataType: "json",
            url: '/delimg',
            data: { ImgName: ImgName },
            success: function (result) {
                if(!result.error) {
                    br.remove();
                    div.remove();
                    return
                }
                alert(result.msg);
            }
        });
    })
}

$("#addimg").click(function(){
    const URL = $("#URL").val();
    const Condition = $("#condition").val();
    if(!URL) {
        alert("URL 不得為空");
        return
    }
    if(!condition) {
        alert("條件 不得為空");
        return
    }
    $.ajax({
        type: "POST",
        dataType: "json",
        url: '/addimg',
        data: { URL: URL , Condition : Condition },
        success: function (result) {
            if(!result.error) {
                alert("新增 " + result.data + " 項");
                getImg();
                return
            }
            alert(result.msg);
        }
    });
})

$("#uploadimg").click(function(){
    let fileData = new FormData();
    fileData.append("file",$("#fileImg")[0].files[0]);
    $.ajax({
        url:"/uploadimg",
        type: "POST",
        data: fileData,
        contentType: false,
        processData: false,
        async: false,
        success: function(result,data){
            if(!result.error) {
                //跳訊息提示
                alert('上傳成功!');
                //清掉假filebox中的內容
                document.getElementById('fileImg').value = '';
                getImg();
                return
            }
            alert(result.msg);
        }
    });
})

$(function () {
    getImg();
});