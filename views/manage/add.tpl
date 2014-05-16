<div class="container">
  <div class="row">
    <div class="hero-text">
     <h1>Add New Article</h1>
     <p class="lead">Add a new article to your account. Once it's successfuly added, you'll be redirected back to your current articles list.</p>
   </div>
 </div>
</div>
</div>
<div class="container">
  <div class="row">
    
    <h2>Article Details</h2>

    {{if .flash.error}}
    <blockquote>{{.flash.error}}</blockquote>
    {{end}}

    {{if .flash.notice}}
    <blockquote>{{.flash.notice}}</blockquote>
    {{end}}
    
    <p>          
      <form role="form" id="user" method="POST">
        <div class="form-group {{if .Errors.Name}}has-error has-feedback{{end}}">
          <label for="name">Article name： {{if .Errors.Name}}({{.Errors.Name}}){{end}}</label>
          <input name="name" type="text" value="{{.Article.Name}}" class="form-control" tabindex="1" />
          {{if .Errors.Name}}<span class="glyphicon glyphicon-remove form-control-feedback"></span>{{end}}
        </div>
        
        <div class="form-group">
          <label for="client">Client：</label>
          <input name="client" type="text" value="{{.Article.Client}}" class="form-control" tabindex="2" />
        </div>
        
        <div class="form-group">
          <label for="url">URL：</label>
          <input name="url" type="text" value="{{.Article.Url}}" class="form-control" tabindex="3" />
        </div>
        
        <input type="submit" value="Create Article" class="btn btn-default" tabindex="4" /> &nbsp;
        <a href="#" title="don't create the article">Cancel</a>
      </form>
    </p>
  </div>
</div>