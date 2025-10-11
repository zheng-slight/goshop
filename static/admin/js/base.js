$(function () {
	baseApp.init();
	//当窗口重置时,重新计算窗口高度
	$(window).resize(function () {
		baseApp.resizeIframe();
	})
})

var baseApp = {
	init: function () {
		this.initAside()
		this.confirmDelete()
		this.resizeIframe()
		this.changeStatus()
		this.changeNum()
	},
	initAside: function () { //左侧菜单栏隐藏显示子栏
		$(".aside h4").click(function () {
			$(this).sibling("ul").slideToggle();
		})
	},
	resizeIframe: function() {  // 设置iframe高度
		$("rightMain").height($(window).height()-80)
	},
	confirmDelete: function () { // 删除提示
		$(".delete").click(function () {
			var flag = confirm("确定要删除该项?")
			return flag
		})
	},
	changeStatus: function () { // 改变状态
		//点击事件
		$(".chStatus").click(function () {
			var $this = $(this);
			var id = $this.attr("data-id");
			var table = $this.attr("data-table");
			var field = $this.attr("data-field");
			console.log(field);
			$.get("/admin/changeStatus", {id: id, table: table, field: field}, function (res) {
				if (res.success) {
					if ($this.attr("src").indexOf("yes") != -1){
						$this.attr("src", "/static/admin/images/no.gif");
					} else {
						$this.attr("src", "/static/admin/images/yes.gif");
					}
				} else {
					alert(res.message);
				}
			})
		})
	},
	changeNum: function () {  // 修改排序数字
		/*
		1、获取el里面的值  var spanNum=$(this).html()
		2、创建一个input的dom节点   var input=$("<input value='' />");
		3、把input放在el里面   $(this).html(input);
		4、让input获取焦点  给input赋值    $(input).trigger('focus').val(val);
		5、点击input的时候阻止冒泡
		$(input).click(function(e){
			e.stopPropagation();
		})
		6、鼠标离开的时候给span赋值,并触发ajax请求
		$(input).blur(function(){
			var inputNum=$(this).val();
			spanEl.html(inputNum);
			触发ajax请求

		})
		*/

		$(".chSpanNum").click(function () {
			// 1、获取el 以及el里面的属性值
			var id = $(this).attr("data-id")
			var table = $(this).attr("data-table")
			var field = $(this).attr("data-field")
			var num = $(this).html().trim()
			var spanEl = $(this)
			//2、创建一个input的dom节点   var input=$("<input value='' />");
			var input = $("<input style='width:60px'  value='' />");
			// 3、把input放在el里面   $(this).html(input);
			$(this).html(input);
			//4、让input获取焦点  给input赋值    $(input).trigger('focus').val(val);
			$(input).trigger("focus").val(num);
			//5、点击input的时候阻止冒泡
			$(input).click(function (e) {
				e.stopPropagation();
			})
			//6、鼠标离开的时候给span赋值,并触发ajax请求
			$(input).blur(function () {
				var inputNum = $(this).val()
				spanEl.html(inputNum)
				//触发ajax请求
				$.get("/admin/changeNum", { id: id, table: table, field: field, num: inputNum }, function (response) {
					if (response.success) {
						alert(response.message)
					} else {
						alert(response.message)
					}
				})
			})
		})

	}
}