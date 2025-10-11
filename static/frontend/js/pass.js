(function ($) {
    $(function () {
        loginApp.init();
    })
    var loginApp = {
        init: function () {
            this.getCaptcha();
            this.captchaImgChage();
            this.initRegisterStep1();
            this.initRegisterStep2();
            this.initRegisterStep3();
            this.initDoLogin();
        },
        getCaptcha: function () { //获取图形验证码
            $.get("/pass/captcha?t=" + Math.random(), function (response) {
                $("#captchaId").val(response.captchaId)
                $("#captchaImg").attr("src", response.captchaImage)
            })
        },
        captchaImgChage: function () { //改变图形验证码
            var _that = this;
            $("#captchaImg").click(function () {
                _that.getCaptcha()
            })
        },
        //注册第一步
        //1.输入手机号以及图形验证码,前端校验是否合法
        //2.点击 立即注册 按钮, 请求后台api,后台会判断手机号以及图形验证码是否符合要求,并发送手机验证码
        initRegisterStep1: function () {
            var _that = this;
            //发送验证码
            $("#registerButton").click(function () {
                //验证验证码是否正确
                var phone = $('#phone').val();
                var verifyCode = $('#verifyCode').val();
                var captchaId = $("#captchaId").val();
                $(".error").html("")
                var reg = /^[\d]{11}$/;
                if (!reg.test(phone)) {
                    $(".error").html("Error：手机号输入错误");
                    return false;
                }
                if (verifyCode.length < 1) {
                    $(".error").html("Error：图形验证码长度不合法")
                    return false;
                }
                //请求后台api,校验输入的手机号以及图形验证码,并发送短信
                $.get("/pass/sendCode", {
                    "phone": phone,
                    "verifyCode": verifyCode,
                    "captchaId": captchaId
                }, function (response) {
                    if (response.success == true) {  //校验成功,进入注册第二步
                        //跳转到下页面
                        location.href = "/pass/registerStep2?sign=" + response.sign + "&verifyCode=" + verifyCode;
                    } else {
                        //改变验证码
                        $(".error").html("Error：" + response.message + ",请重新输入!")
                        //改变验证码
                        _that.getCaptcha()
                    }
                })

            })
        },
        // 注册第二步
        //输入短信验证码,校验,成功则跳转到注册第三步
        //还可以重新发送短信验证码
        initRegisterStep2: function () {
            $(function () {
                var timer = 10;

                function Countdown() {  // 重新发送短信时间
                    if (timer >= 1) {
                        timer -= 1;
                        $("#sendCode").attr('disabled', true);
                        $("#sendCode").html('重新发送(' + timer + ')');
                        setTimeout(function () {
                            Countdown();
                        }, 1000);
                    } else {
                        $("#sendCode").attr('disabled', false)
                        $("#sendCode").html('重新发送');
                    }
                }

                Countdown();
                //重新发送短信
                $("#sendCode").click(function () {
                    timer = 10;
                    Countdown();
                    var phone = $("#phone").val()
                    var verifyCode = $("#verifyCode").val()
                    var captchaId = "resend"  //重新发送标签

                    //重新请求接口发送短信
                    $.get("/pass/sendCode", {
                        "phone": phone,
                        "verifyCode": verifyCode,
                        "captchaId": captchaId
                    }, function (response) {
                        console.log(response)
                    })
                })
            })

            //验证验证码
            $(function () {
                $("#nextStep").click(function (e) {
                    $(".error").html()
                    var sign = $('#sign').val();
                    var smsCode = $('#smsCode').val();
                    //请求api,校验输入短信是否正确,并跳转到注册第三步
                    $.get('/pass/validateSmsCode', {sign: sign, smsCode: smsCode}, function (response) {
                        if (response.success == true) {
                            location.href = "/pass/registerStep3?sign=" + sign + "&smsCode=" + smsCode
                        } else {
                            $(".error").html("Error：" + response.message)
                        }
                    })
                })

                $("#returnButton").click(function () {
                    location.href = "/pass/registerStep1"
                })
            })
        },
        initRegisterStep3: function () { //注册第三步: 设置密码注册用户
            $(function () {
                $("#form").submit(function () {
                    $(".error").html("")
                    var password = $('#password').val();
                    var rpassword = $('#rpassword').val();

                    if (password.length < 6) {
                        $(".error").html("Error：密码的长度不能小于6位")
                        return false;
                    }
                    if (password != rpassword) {
                        $(".error").html("Error：密码和确认密码不一致")
                        return false;
                    }
                    return true;

                })
            })
        },
        initDoLogin: function () { //登录操作
            var _that = this;
            $("#doLogin").click(function (e) {
                $(".error").html("")
                var phone = $('#phone').val();
                var password = $('#password').val();
                var captchaId = $('#captchaId').val();
                var captchaVal = $("#captchaVal").val();
                //获取返回上一页的地址
                var prevPage = $("#prevPage").val();

                var reg = /^[\d]{11}$/;
                if (!reg.test(phone)) {
                    $(".error").html('Error:手机号输入错误');
                    return false;
                }
                if (password.length < 6) {
                    $(".error").html('Error:密码长度不合法');
                    return false;
                }

                if (captchaVal.length < 1) {
                    $(".error").html('Error:验证码长度不合法');
                    return false;
                }
                //ajax请求
                $.post('/pass/doLogin', {
                    phone: phone,
                    password: password,
                    captchaVal: captchaVal,
                    captchaId: captchaId
                }, function (response) {  //登录成功跳转到首页
                    if (response.success == true) {
                        if (prevPage == "") {
                            location.href = "/";
                        } else {
                            location.href = prevPage;
                        }
                    } else {
                        $(".error").html("Error：" + response.message + ",请重新输入!")
                        //改变验证码
                        _that.getCaptcha();
                    }
                })
            })
        }
    }
})($)

