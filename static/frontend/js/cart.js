(function ($) {
    var app = {
        init: function () {
            this.changeCartNum();
            this.deleteConfirm();
            this.initCheckBox();
            this.isCheckedAll();
            this.initChekOut();
        },
        initChekOut(){ //点击'去结算'按钮时,判断是否选中购物车,如果选中了,则跳转到 '确认订单' 页面
            $(function(){
                $("#checkout").click(function(){
                    var allPrice=parseFloat($("#allPrice").html());
                    if(allPrice==0){
                        alert('购物车没有选中去结算的商品')
                    }else{
                        location.href="/buy/checkout";
                    }
                })
            })
        },
        deleteConfirm: function () { //删除确认
            $('.delete').click(function () {
                var flag = confirm('您确定要删除吗?');
                return flag;
            })

        },
        changeCartNum() {  // 改变购物车数据
            $('.decCart').click(function () {  //减少数量
                //获取商品id,商品颜色
                var goods_id = $(this).attr("goods_id")
                var goods_color = $(this).attr("goods_color")
                var _that = this;
                //请求url
                $.get('/cart/decCart?goods_id=' + goods_id + '&goods_color=' + goods_color, function (response) {
                    if (response.success) { //改变成功
                        //更新总商品价格
                        $("#allPrice").html(response.allPrice + "元")
                        //更新对应商品数量
                        $(_that).siblings(".input_center").find("input").val(response.num)
                        //更新对应商品价格小计
                        $(_that).parent().parent().siblings(".totalPrice").html(response.currentPrice + "元")
                    }
                })
            });

            $('.incCart').click(function () {  // 增加数量
                var goods_id = $(this).attr("goods_id")
                var goods_color = $(this).attr("goods_color")
                var _that = this;
                $.get('/cart/incCart?goods_id=' + goods_id + '&goods_color=' + goods_color, function (response) {
                    console.log(response)
                    if (response.success) {
                        $("#allPrice").html(response.allPrice + "元")
                        $(_that).siblings(".input_center").find("input").val(response.num)
                        $(_that).parent().parent().siblings(".totalPrice").html(response.currentPrice + "元")
                    }
                })
            });
        },

        initCheckBox(){  //改变一个商品数据的选中状态
            //全选按钮点击
            $("#checkAll").click(function() {
                if (this.checked) {
                    $(":checkbox").prop("checked", true);
                    //让cookie中商品的checked属性都等于true
                    $.get('/cart/changeAllCart?flag=1',function(response){
                        if(response.success){
                            $("#allPrice").html(response.allPrice+"元")
                        }
                    })
                }else {
                    $(":checkbox").prop("checked", false);
                    //让cookie中商品的checked属性都等于false
                    $.get('/cart/changeAllCart?flag=0',function(response){
                        if(response.success){
                            $("#allPrice").html(response.allPrice+"元")
                        }
                    })
                }
            });

            //点击单个选择框按钮的时候触发
            var _that=this;
            $(".cart_list :checkbox").click(function() {
                _that.isCheckedAll();
                var goods_id=$(this).attr("goods_id")
                var goods_color=$(this).attr("goods_color")
                $.get('/cart/changeOneCart?goods_id='+goods_id+'&goods_color='+goods_color,function(response){
                    if(response.success){
                        $("#allPrice").html(response.allPrice+"元")
                    }
                })


            });   //注意：this指向
        },

        isCheckedAll(){   //判断全选是否选择
            var allNum = $(".cart_list :checkbox").size();//checkbox总个数
            var checkedNum = 0;
            $(".cart_list :checkbox").each(function () {
                if($(this).prop("checked")==true){
                    checkedNum++;
                }
            });
            if(allNum==checkedNum){//全选
                $("#checkAll").prop("checked",true);
            }else{//不全选
                $("#checkAll").prop("checked",false);
            }
        },
    }

    $(function () {
        app.init();
    })
})($)
