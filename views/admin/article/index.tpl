        <section class="content">            <table class="table">                <caption>文章列表</caption>                <thead>                <tr>                    <th>#</th>                    <th>文章标题</th>                    <th>创建时间</th>                    <th>作者</th>                    <th>编辑</th>                </tr>                </thead>                <tbody>{{range $key, $val := .articles}}                <tr>                    <td>{{$val.Id}}</td>                    <td>{{$val.Title}}</td>                    <td>{{date  $val.CreatedAt "Y-m-d H:i:s"}}</td>                    <td>@twitter</td>                    <td><a href="edit">编辑</a></td>                </tr>{{end}}                </tbody>            </table>