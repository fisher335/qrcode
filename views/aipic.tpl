{{template "aa.tpl"}}
<div class="container-fluid">
	<div class="row">
		<div class="col-md-12">
				<div class="text-center" style="margin:20px">
					<img  src=" {{.path}} " class="center-block" height="300" width="300" />
				</div>
            <div class="divider"></div>
			<div class="jumbotron text-center">
				<h2>
					二维码结果
				</h2>
				<p>
					生成二维码如图，请右键保存，我就不加按钮了。
				</p>
				<p>
					<a class="btn btn-primary btn-large" href=" {{.path}}">下载图片</a>
				</p>
			</div>
		</div>
	</div>
</div>
{{template "zz.tpl"}}