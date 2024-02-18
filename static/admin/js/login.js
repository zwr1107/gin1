$(function(){
    app.init();
})
var app={
    init:function(){
        this.getCaptcha()
        this.captchaImgChage()
    },
    getCaptcha:function(){
        $.get("/admin/getCaptcha?t="+Math.random(),function(response){
            console.log(response)
            $("#captchaId").val(response.data.id)
            $("#captchaImg").attr("src",response.data.image)
        })
    },
    captchaImgChage:function(){
        var that=this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}