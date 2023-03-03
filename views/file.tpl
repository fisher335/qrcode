{{template "aa.tpl"}}

<table style = " margin-top:5px; margin-left:50px;margin-right:50px;"  class="table table-bordered "    >
	<caption>视频文件</caption>

   <thead>
      <tr>
         <th>文件名称</th>
          <th>操作</th>
      </tr>
   </thead>
   <tbody>
  {{range $file:=  .file}}

          <tr>
                <td>{{$file}}</td>
             <td><a href="/download/videos/{{$file}}">播放</a></td>
          </tr>


{{end}}
   </tbody>
</table>
{{template "zz.tpl"}}