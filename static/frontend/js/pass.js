(function ($) {
    "use strict";

    var COUNTDOWN_SECONDS = 10;

    function setLoading($btn, loading) {
        if (!$btn || !$btn.length) {
            return;
        }
        if (loading) {
            $btn.data("originalText", $btn.text());
            $btn.prop("disabled", true).text("处理中...");
        } else {
            var originalText = $btn.data("originalText");
            if (originalText) {
                $btn.text(originalText);
            }
            $btn.prop("disabled", false);
        }
    }

    function setFieldError(id, message) {
        var $input = $("#" + id);
        var $error = $("#" + id + "-error");
        if (message) {
            if ($input.length) {
                $input.attr("aria-invalid", "true");
            }
            if ($error.length) {
                $error.text(message);
            }
        } else {
            if ($input.length) {
                $input.removeAttr("aria-invalid");
            }
            if ($error.length) {
                $error.text("");
            }
        }
    }

    function setFormError($form, message) {
        var $error = $form.find("#form-error");
        if ($error.length) {
            $error.text(message || "");
        }
    }

    function clearErrors($form) {
        $form.find("[aria-invalid]").removeAttr("aria-invalid");
        $form.find(".field-error").text("");
        setFormError($form, "");
    }

    var loginApp = {
        init: function () {
            this.initCaptcha();
            this.initLogin();
            this.initRegisterStep1();
            this.initRegisterStep2();
            this.initRegisterStep3();
        },

        getCaptcha: function () {
            var $captchaId = $("#captchaId");
            var $captchaImg = $("#captchaImg");
            if (!$captchaId.length || !$captchaImg.length) {
                return;
            }
            $.get("/pass/captcha?t=" + Math.random(), function (response) {
                $captchaId.val(response.captchaId);
                $captchaImg.attr("src", response.captchaImage);
            });
        },

        initCaptcha: function () {
            var _that = this;
            $("#captchaImg").click(function () {
                _that.getCaptcha();
            });
            this.getCaptcha();
        },

        initLogin: function () {
            var _that = this;
            $("#loginForm").submit(function (event) {
                event.preventDefault();
                var $form = $(this);
                clearErrors($form);

                var phone = $.trim($("#phone").val());
                var password = $("#password").val();
                var captchaVal = $("#captchaVal").val();
                var captchaId = $("#captchaId").val();
                var prevPage = $("#prevPage").val();

                if (!/^[\d]{11}$/.test(phone)) {
                    setFieldError("phone", "手机号输入错误");
                    $("#phone").focus();
                    return;
                }
                if (password.length < 6) {
                    setFieldError("password", "密码长度不合法");
                    $("#password").focus();
                    return;
                }
                if (!captchaVal) {
                    setFieldError("captchaVal", "验证码不能为空");
                    $("#captchaVal").focus();
                    return;
                }

                var $btn = $("#doLogin");
                setLoading($btn, true);
                $.post("/pass/doLogin", {
                    phone: phone,
                    password: password,
                    captchaVal: captchaVal,
                    captchaId: captchaId
                }, function (response) {
                    setLoading($btn, false);
                    if (response.success === true) {
                        location.href = prevPage || "/";
                    } else {
                        setFormError($form, response.message + "，请重新输入");
                        _that.getCaptcha();
                        $("#captchaVal").val("");
                    }
                }).fail(function () {
                    setLoading($btn, false);
                    setFormError($form, "网络异常，请稍后重试");
                });
            });
        },

        initRegisterStep1: function () {
            var _that = this;
            $("#registerForm").submit(function (event) {
                event.preventDefault();
                var $form = $(this);
                clearErrors($form);

                var phone = $.trim($("#phone").val());
                var verifyCode = $("#verifyCode").val();
                var captchaId = $("#captchaId").val();

                if (!/^[\d]{11}$/.test(phone)) {
                    setFieldError("phone", "手机号输入错误");
                    $("#phone").focus();
                    return;
                }
                if (!verifyCode) {
                    setFieldError("verifyCode", "图形验证码不能为空");
                    $("#verifyCode").focus();
                    return;
                }

                var $btn = $("#registerButton");
                setLoading($btn, true);
                $.get("/pass/sendCode", {
                    phone: phone,
                    verifyCode: verifyCode,
                    captchaId: captchaId
                }, function (response) {
                    setLoading($btn, false);
                    if (response.success === true) {
                        location.href = "/pass/registerStep2?sign=" + encodeURIComponent(response.sign) + "&verifyCode=" + encodeURIComponent(verifyCode);
                    } else {
                        setFormError($form, response.message + "，请重新输入");
                        _that.getCaptcha();
                        $("#verifyCode").val("");
                    }
                }).fail(function () {
                    setLoading($btn, false);
                    setFormError($form, "网络异常，请稍后重试");
                });
            });
        },

        initRegisterStep2: function () {
            var timer = COUNTDOWN_SECONDS;
            var countdownTimer = null;
            var $sendCode = $("#sendCode");

            function renderCountdown() {
                if (timer >= 1) {
                    $sendCode.prop("disabled", true).text("重新发送(" + timer + ")");
                } else {
                    $sendCode.prop("disabled", false).text("重新发送");
                }
            }

            function startCountdown() {
                if (countdownTimer) {
                    clearInterval(countdownTimer);
                }
                timer = COUNTDOWN_SECONDS;
                renderCountdown();
                countdownTimer = setInterval(function () {
                    timer -= 1;
                    renderCountdown();
                    if (timer <= 0) {
                        clearInterval(countdownTimer);
                        countdownTimer = null;
                    }
                }, 1000);
            }

            $sendCode.click(function () {
                startCountdown();
                $.get("/pass/sendCode", {
                    phone: $("#phone").val(),
                    verifyCode: $("#verifyCode").val(),
                    captchaId: "resend"
                });
            });

            $("#smsForm").submit(function (event) {
                event.preventDefault();
                var $form = $(this);
                clearErrors($form);

                var sign = $("#sign").val();
                var smsCode = $("#smsCode").val();

                if (!smsCode) {
                    setFieldError("smsCode", "请输入短信验证码");
                    $("#smsCode").focus();
                    return;
                }

                var $btn = $("#nextStep");
                setLoading($btn, true);
                $.get("/pass/validateSmsCode", { sign: sign, smsCode: smsCode }, function (response) {
                    setLoading($btn, false);
                    if (response.success === true) {
                        location.href = "/pass/registerStep3?sign=" + encodeURIComponent(sign) + "&smsCode=" + encodeURIComponent(smsCode);
                    } else {
                        setFormError($form, response.message);
                    }
                }).fail(function () {
                    setLoading($btn, false);
                    setFormError($form, "网络异常，请稍后重试");
                });
            });

            $("#returnButton").click(function () {
                location.href = "/pass/registerStep1";
            });

            startCountdown();
        },

        initRegisterStep3: function () {
            $("#form").submit(function () {
                clearErrors($(this));

                var password = $("#password").val();
                var rpassword = $("#rpassword").val();

                if (password.length < 6) {
                    setFieldError("password", "密码的长度不能小于6位");
                    $("#password").focus();
                    return false;
                }
                if (password !== rpassword) {
                    setFieldError("rpassword", "两次输入的密码不一致");
                    $("#rpassword").focus();
                    return false;
                }
                return true;
            });
        }
    };

    $(function () {
        loginApp.init();
    });
})(jQuery);
