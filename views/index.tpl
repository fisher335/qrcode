{{template "aa.tpl"}}
<div class="container-fluid">
    <div class="row">
        <div class="col-md-12">
            <form role="form" action="/qrcode/" method="post" id="url_form">
                <div class="form-group">

                    <label for="exampleInputEmail1">
                        二维码地址
                    </label>
                    <input class="form-control" id="exampleInputEmail1" name="url"/>

                    <p class="help-block">
                        输入你想生成二维码的地址
                    </p>
                    <div class="checkbox">

                        <label>
                            <input type="checkbox"/> 确认提交
                        </label>
                    </div>
                    <button type="button" class="btn btn-primary" onclick=sub()>
                        提交
                    </button>
                </div>
            </form>
        </div>
    </div>
</div>
<script>
    function sub() {
        if ($("#Email").val().trim() == "") {
            alert("请输入值再提交");
            return;
        } else {
            $("#url_form").submit();
        }

    }

</script>
{{template "zz.tpl"}}