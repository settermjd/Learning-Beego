<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Current Article</h1>
     <p class="lead">Here you'll find all the available articles</p>
     {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    <table class="table">
      <thead>
        <tr>
          <th>Id</th>
          <th>Name</th>
          <th>Client</th>
          <th>Url</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td>{{$record.Id}}</td>
          <td>{{$record.Name}}</td>
          <td>{{$record.Client}}</td>
          <td>{{$record.Url}} {{urlfor "ManageController.Delete" ":id" "21"}}</td>
        </tr>
        {{end}}
      </tbody>
      <tfoot>
        <tr>
          <td colspan="4"><a href="{{urlfor "ManageController.Add"}}" title="add new article">add new article</a></td>
        </tr>
      </tfoot>
    </table>
  </div>
</div>