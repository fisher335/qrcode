{{template "aa.tpl"}}
<table class="table table-bordered" align="center" sortable="true">
    <caption>信息查询</caption>

    <thead>
    <tr>
        <th>键</th>
        <th>值</th>
    </tr>
    </thead>
    <tbody>

    {{range $key, $val :=  .head}}
    <tr>
        <td>{{$key}}</td>
        <td>{{$val}}</td>
    </tr>
    {{end}}
    </tbody>
</table>
{{template "zz.tpl"}}