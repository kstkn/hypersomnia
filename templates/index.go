// Code generated by go generate; DO NOT EDIT.
package templates

const Index = `<!doctype html><html lang=en><meta charset=utf-8><meta name=viewport content="width=device-width,initial-scale=1,shrink-to-fit=no"><link rel=dns-prefetch href=//unpkg.com><link href="data:image/x-icon;base64,AAABAAEAEBAAAAEACABoBQAAFgAAACgAAAAQAAAAIAAAAAEACAAAAAAAAAEAAAAAAAAAAAAAAAEAAAAAAAAAAAAAyMnMAJtFOACTOi4AyMvPAFRF0QDAyc0A0djcAJOEgABxcnIAmD8xAMjP1QCzvMEAkqPPADAZZwB+f38AMhMMAHGm4QDH0dgAY0I8AFRAOgDr7O4Ahi4gAJM5KgBzM00AbSMXAHcnHADKOSgAMhQKAIYyJgBzLB8Ae3t7AGtgyABIGQwAhTQpANLY3gCBMCQAjTUmAHwsHwC3u70Ar5GOAHJ1dACKMCQAusPIAFw7pQDCxMUAaj0zACgpLACQOC8AdSgbAG1JQwDGzNAAfYB/AH0uIwCSlZYAytLYAHRlYADyW0UAhi8gAHFzcwBHYmsAMxYgAMfNzgCFMyYAkTgoAHeM2wCHMyYAfSgcAMhnYADAxMcAbiQYAI00LACBLCIAfykaAMjP0gAwEwkAkzosAMfQ2AAzFQwAW1OGADUVDAB2ea8APCEcAIArHgDcQS8AsZmVAFxZfgAtFQ0AY2CDADIVDQB+gYAAdXZ2AJeanAAqEmkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAXDITNgAAAAAAAAAAAAAAOC4oVQgUAAAAAAAAAAAAAB8GTB0eVwAAAAAAAAAAAAAPCwolJFkAAAAAAAAAAAAANDcCQD9OAAAAAAAAAAAAAFpNIxdCUAAAAAAAAAAAAABPEgcDIj0AAAAAAAAAAAAAUlFBEQ1bAAAAAAAAAAAAAA4YMAUgWAAAAAAAAAAAAABdNUcsSlYAAAAAAAAAAAAAHCYqSDMpAAAAAAAAAAAAABAxOhY+OwAAAAAAAAAAAABLRlNJBAkAAAAAAAAAAAAALxlDGgE7AAAAAAAAAAAAACEbOVREPAAAAAAAAAAAAAAMJy0VRSsAAAAAAPw/AAD4HwAA+B8AAPgfAAD4HwAA+B8AAPgfAAD4HwAA+B8AAPgfAAD4HwAA+B8AAPgfAAD4HwAA+B8AAPgfAAA=" rel=icon type=image/x-icon><link rel=stylesheet href=https://unpkg.com/bootstrap@4.3.1/dist/css/bootstrap.min.css><link rel=stylesheet href=https://unpkg.com/jsoneditor@6.1.0/dist/jsoneditor.min.css><title>Hypersomnia</title><style>body,html{margin:0;height:100%}body{overflow:hidden}.symbol{display:none}.icon{width:1em;height:1em;vertical-align:-.125em;fill:#212529}.icon-danger{fill:#dc3545!important}.container-fluid,.parent{height:100%}#services,#request,#response{position:relative;float:left;height:100%;overflow-y:auto}div.active{background-color:#444;color:#fff;-webkit-border-radius:3px;-moz-border-radius:3px;border-radius:3px}div.active .icon{fill:#fff}.ace_editor{font:10px roboto-mono,Monaco,Menlo,ubuntu mono,Consolas,monospace!important}#request-body-editor{height:70%}#request-context-editor{height:15%}#request-context-editor textarea{width:100%;height:100%;font:10px roboto-mono,Monaco,Menlo,ubuntu mono,Consolas,monospace!important}#response-editor{height:90%}.clickable{cursor:pointer}</style><svg class="symbol"><symbol id="cog" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M487.4 315.7l-42.6-24.6c4.3-23.2 4.3-47 0-70.2l42.6-24.6c4.9-2.8 7.1-8.6 5.5-14-11.1-35.6-30-67.8-54.7-94.6-3.8-4.1-10-5.1-14.8-2.3L380.8 110c-17.9-15.4-38.5-27.3-60.8-35.1V25.8c0-5.6-3.9-10.5-9.4-11.7-36.7-8.2-74.3-7.8-109.2.0-5.5 1.2-9.4 6.1-9.4 11.7V75c-22.2 7.9-42.8 19.8-60.8 35.1L88.7 85.5c-4.9-2.8-11-1.9-14.8 2.3-24.7 26.7-43.6 58.9-54.7 94.6-1.7 5.4.6 11.2 5.5 14L67.3 221c-4.3 23.2-4.3 47 0 70.2l-42.6 24.6c-4.9 2.8-7.1 8.6-5.5 14 11.1 35.6 30 67.8 54.7 94.6 3.8 4.1 10 5.1 14.8 2.3l42.6-24.6c17.9 15.4 38.5 27.3 60.8 35.1v49.2c0 5.6 3.9 10.5 9.4 11.7 36.7 8.2 74.3 7.8 109.2.0 5.5-1.2 9.4-6.1 9.4-11.7v-49.2c22.2-7.9 42.8-19.8 60.8-35.1l42.6 24.6c4.9 2.8 11 1.9 14.8-2.3 24.7-26.7 43.6-58.9 54.7-94.6 1.5-5.5-.7-11.3-5.6-14.1zM256 336c-44.1.0-80-35.9-80-80s35.9-80 80-80 80 35.9 80 80-35.9 80-80 80z"/></symbol></svg><svg class="symbol"><symbol id="cube" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><path d="M239.1 6.3l-208 78c-18.7 7-31.1 25-31.1 45v225.1c0 18.2 10.3 34.8 26.5 42.9l208 104c13.5 6.8 29.4 6.8 42.9.0l208-104c16.3-8.1 26.5-24.8 26.5-42.9V129.3c0-20-12.4-37.9-31.1-44.9l-208-78C262 2.2 250 2.2 239.1 6.3zM256 68.4l192 72v1.1l-192 78-192-78v-1.1l192-72zm32 356V275.5l160-65v133.9l-160 80z"/></symbol></svg><svg class="symbol"><symbol id="star-solid" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path d="M259.3 17.8 194 150.2 47.9 171.5c-26.2 3.8-36.7 36.1-17.7 54.6l105.7 103-25 145.5c-4.5 26.3 23.2 46 46.4 33.7L288 439.6l130.7 68.7c23.2 12.2 50.9-7.4 46.4-33.7l-25-145.5 105.7-103c19-18.5 8.5-50.8-17.7-54.6L382 150.2 316.7 17.8c-11.7-23.6-45.6-23.9-57.4.0z"/></symbol></svg><svg class="symbol"><symbol id="star-regular" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path d="M528.1 171.5 382 150.2 316.7 17.8c-11.7-23.6-45.6-23.9-57.4.0L194 150.2 47.9 171.5c-26.2 3.8-36.7 36.1-17.7 54.6l105.7 103-25 145.5c-4.5 26.3 23.2 46 46.4 33.7L288 439.6l130.7 68.7c23.2 12.2 50.9-7.4 46.4-33.7l-25-145.5 105.7-103c19-18.5 8.5-50.8-17.7-54.6zM388.6 312.3l23.7 138.4L288 385.4l-124.3 65.3 23.7-138.4-100.6-98 139-20.2 62.2-126 62.2 126 139 20.2-100.6 98z"/></symbol></svg><svg class="symbol"><symbol id="paste-solid" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M128 184c0-30.879 25.122-56 56-56h136V56c0-13.255-10.745-24-24-24h-80.61C204.306 12.89 183.637.0 160 0s-44.306 12.89-55.39 32H24C10.745 32 0 42.745.0 56v336c0 13.255 10.745 24 24 24h104V184zm32-144c13.255.0 24 10.745 24 24s-10.745 24-24 24-24-10.745-24-24 10.745-24 24-24zm184 248h104v2e2c0 13.255-10.745 24-24 24H184c-13.255.0-24-10.745-24-24V184c0-13.255 10.745-24 24-24h136v104c0 13.2 10.8 24 24 24zm104-38.059V256h-96v-96h6.059a24 24 0 0116.97 7.029l65.941 65.941a24.002 24.002.0 017.03 16.971z"/></symbol></svg><svg class="symbol"><symbol id="edit-solid" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"><path d="M402.6 83.2l90.2 90.2c3.8 3.8 3.8 10 0 13.8L274.4 405.6l-92.8 10.3c-12.4 1.4-22.9-9.1-21.5-21.5l10.3-92.8L388.8 83.2c3.8-3.8 10-3.8 13.8.0zm162-22.9-48.8-48.8c-15.2-15.2-39.9-15.2-55.2.0l-35.4 35.4c-3.8 3.8-3.8 10 0 13.8l90.2 90.2c3.8 3.8 10 3.8 13.8.0l35.4-35.4c15.2-15.3 15.2-40 0-55.2zM384 346.2V448H64V128h229.8c3.2.0 6.2-1.3 8.5-3.5l40-40c7.6-7.6 2.2-20.5-8.5-20.5H48C21.5 64 0 85.5.0 112v352c0 26.5 21.5 48 48 48h352c26.5.0 48-21.5 48-48V306.2c0-10.7-12.9-16-20.5-8.5l-40 40c-2.2 2.3-3.5 5.3-3.5 8.5z"/></symbol></svg><svg class="symbol"><symbol id="trash-solid" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512"><path d="M432 32H312l-9.4-18.7A24 24 0 00281.1.0H166.8a23.72 23.72.0 00-21.4 13.3L136 32H16A16 16 0 000 48v32a16 16 0 0016 16h416a16 16 0 0016-16V48a16 16 0 00-16-16zM53.2 467a48 48 0 0047.9 45h245.8a48 48 0 0047.9-45L416 128H32z"/></symbol></svg><div class="container-fluid p-2 pl-4"><div class="row parent"><div class=col-sm id=services style=overflow-y:hidden>{{ $length := len .Envs }}<div {{if le $length 1}} style=display:none {{end}}><label for=environmentSwitcher>Environment:</label> <select id=environmentSwitcher class=js-environment>
{{range .Envs}}<option value={{.}}>{{.}}{{end}}</select><svg class="icon clickable js-environment-settings ml-1"><use href="#cog"/></svg></div><div class="mt-1 mb-1">Services:</div><div class=js-services style=overflow-y:auto;height:90%></div></div><div class=col-sm id=request><div class="mb-2 clearfix"><div class="js-active-endpoint float-left pt-1 pb-1"><span class=text-muted>Select endpoint on the left</span></div><button class="btn btn-sm btn-primary js-send float-right">Send</button>
<button class="btn btn-sm btn-secondary js-reset float-right mr-2">Reset</button></div><div id=request-body-editor></div><span class=text-muted>Context (one per line, <code>:</code> - delimited)</span><div id=request-context-editor><textarea></textarea></div></div><div class=col-sm id=response><div class="mb-2 mt-3 clearfix"><span class="badge badge-secondary js-response-time float-right">...</span>
<span class="badge badge-secondary js-response-took float-right mr-2">...</span>
<span class="badge badge-secondary js-correlation-id float-right mr-2" style=cursor:pointer>...</span></div><div id=response-editor></div></div></div></div><div class="modal fade" id=environmentSettings tabindex=-1 role=dialog><div class=modal-dialog role=document><div class=modal-content><div class=modal-header><h5 class=modal-title>Environment settings</h5><button type=button class=close data-dismiss=modal aria-label=Close>
<span aria-hidden=true>&#215;</span></button></div><div class=modal-body><h5>External logs</h5><p>Enter external log systems URLs here using <code>_correlationId_</code> as placeholer.
When you click correlation id badge on top of response box the placeholders in URL will be
substituted with actual value and URL will open.<div class=js-environment-logs-urls></div></div><div class=modal-footer><button type=button class="btn btn-primary js-environment-settings-save" data-dismiss=modal>Save</button>
<button type=button class="btn btn-secondary" data-dismiss=modal>Close</button></div></div></div></div><div class="modal fade" id=newRequest tabindex=-1 role=dialog><div class=modal-dialog role=document><div class=modal-content><div class=modal-header><h5 class=modal-title>New request</h5><button type=button class=close data-dismiss=modal aria-label=Close>
<span aria-hidden=true>&#215;</span></button></div><div class=modal-body><p class=js-new-request-endpoint><div class=form-group><label for=newRequestName>Name</label>
<input class="form-control form-control-sm js-new-request-name" id=newRequestName></div></div><div class=modal-footer><button type=button class="btn btn-primary js-new-request-save" data-dismiss=modal>Save</button>
<button type=button class="btn btn-secondary" data-dismiss=modal>Close</button></div></div></div></div><script src=https://unpkg.com/jquery@3.3.1/dist/jquery.min.js></script><script src=https://unpkg.com/popper.js@1.15.0/dist/umd/popper.min.js></script><script src=https://unpkg.com/bootstrap@4.3.1/dist/js/bootstrap.min.js></script><script src=https://unpkg.com/moment@2.24.0/min/moment.min.js></script><script src=https://unpkg.com/jsoneditor@6.1.0/dist/jsoneditor.min.js></script><script src=https://unpkg.com/dexie@2.0.4/dist/dexie.js></script><script src=https://unpkg.com/jsonpath@1.0.2/jsonpath.min.js></script><script src=https://unpkg.com/mustache@3.0.1/mustache.min.js></script>{{ .JsTemplates }}
<script>let tmplService=$('#tmplService').html();let tmplRequest=$('#tmplRequest').html();let tmplEnvironmentLogsUrl=$('#tmplEnvironmentLogsUrl').html();String.prototype.replaceAll=function(search,replacement){let target=this;return target.replace(new RegExp(search,'g'),replacement);};Array.prototype.remove=function(){let what,a=arguments,L=a.length,ax;while(L&&this.length){what=a[--L];while((ax=this.indexOf(what))!==-1){this.splice(ax,1);}}
return this;};String.prototype.splitN=function(separator,limit){let str=this;str=str.split(separator);if(str.length>limit){let ret=str.splice(0,limit);ret.push(str.join(separator));return ret;}
return str;};let copyToClipboard=function(data){var tempInput=document.createElement('INPUT');let body=$('body')[0];body.appendChild(tempInput);tempInput.setAttribute('value',data);tempInput.select();document.execCommand('copy');body.removeChild(tempInput);};function pregQuote(str,delimiter){return(str+'').replace(new RegExp('[.\\\\+*?\\[\\^\\]$(){}=!<>|:\\'+(delimiter||'/')+'-]','g'),'\\$&');}
function getMatches(string){let matches=[];let match;let regex=/Response\((.+?),(.+?),(.+?),?(\s?int)?\)/g;while(match=regex.exec(string)){matches[match[0]]={service:match[1],request:match[2],path:match[3],type:match[4]||'string',};}
return matches;}
function pack(v){let m={};if(v!==null&&typeof v!=='undefined'&&typeof v.values!=='undefined'){for(let i in v.values){if(v.values[i].name===''||v.values[i].name==='-'){continue;}
if(v.values[i].values==null||v.values[i].values.length==0){if(v.values[i].type=='string'){m[v.values[i].name]="";}else{m[v.values[i].name]=0;}}else{m[v.values[i].name]=pack(v.values[i]);}}}
return m}
let storage=window.localStorage;Dexie.debug=false;let db=new Dexie("hypersomnia");db.version(1).stores({requests:'name,service,endpoint,body,static',responses:'[service+request],time,receivedAt,body,correlationId',});let activeEnvironment=storage.getItem('active-environment');let activeRequest=storage.getItem('active-request');let favorites=storage.getItem('favorites');if(favorites){favorites=JSON.parse(favorites);}else{favorites=[];}
$(function(){let requestBodyEditor=new JSONEditor(document.getElementById("request-body-editor"),{enableSort:false,enableTransform:false,mode:'code',});let responseEditor=new JSONEditor(document.getElementById("response-editor"),{mode:'code',});let environmentSwitcher=$('.js-environment');environmentSwitcher.on('change',function(){let environment=$(this).val();$(this).prop('disabled',true);$.ajax({method:'POST',url:'services',dataType:'json',contentType:'application/json',data:JSON.stringify({environment:environment,}),success:function(responseServices){localStorage.setItem('active-environment',environment);let services=$('.js-services');services.text('');for(let i in responseServices){let name=responseServices[i].name;if(favorites.indexOf(name)===-1){continue;}
let id=name.replace(/\./g,'_');services.append(Mustache.render(tmplService,{id:id,name:name}));}
for(let i in responseServices){let name=responseServices[i].name;if(name===''){continue;}
if(favorites.indexOf(name)!==-1){continue;}
let id=name.replace(/\./g,'_');services.append(Mustache.render(tmplService,{id:id,name:name}));}
for(let i in favorites){$('.js-favorite[data-service="'+favorites[i]+'"]').addClass('js-favorite-on').find('use').attr('href','#star-solid');}
$('.collapse').each(function(){let id=$(this).attr('id');if(storage.getItem(id+':show')==='true'){$(this).collapse('show');}else{$(this).collapse('hide');}});},error:function(jqXHR,_,errorThrown){alert(errorThrown+': '+jqXHR.responseText);},complete:function(){$('.js-environment').prop('disabled',false)}});});$(document).on('show.bs.collapse','.collapse',function(){let id=$(this).attr('id');storage.setItem(id+':show','true');let service=$(this).data('service');let container=$('.js-requests[id="'+id+'"]');container.html('loading...');class ResponseService{endpoints;version;}
$.ajax({method:'POST',url:'service',dataType:'json',contentType:'application/json',data:JSON.stringify({environment:$('.js-environment').val(),name:service,}),success:function(responseService){$('.js-service[data-service="'+service+'"] .js-version').text(responseService.version);container.html('<ul class="list-unstyled"></ul>');let requests=[];db.requests.where({service:service,static:true}).each(function(request){requests.push(request)}).then(function(){for(let i in responseService.endpoints){if(!responseService.endpoints.hasOwnProperty(i)){continue;}
if(responseService.endpoints[i].name===''){continue;}
responseService.endpoints[i].endpoint=responseService.endpoints[i].name;responseService.endpoints[i].static=false;requests.push(responseService.endpoints[i]);}
for(let i in requests){if(!requests.hasOwnProperty(i)){continue;}
container.find('ul').append(Mustache.render(tmplRequest,{name:requests[i].name,service:service,endpoint:requests[i].endpoint,requestBodyTemplate:requests[i].request?JSON.stringify(pack(requests[i].request)):'{}',static:requests[i].static,}));}
if(activeRequest){$('.js-request[data-name="'+activeRequest+'"]').trigger('click');}});},error:function(jqXHR,_,errorThrown){alert(errorThrown+': '+jqXHR.responseText);},complete:function(){}});}).on('hide.bs.collapse','.collapse',function(){let id=$(this).attr('id');storage.removeItem(id+':show');});$('.collapse').each(function(){let id=$(this).attr('id');if(storage.getItem(id+':show')==='true'){$(this).collapse('show');}else{$(this).collapse('hide');}});$(document).on('click','.js-request',function(){let service=$(this).data('service');storage.setItem('active-service',service);let name=$(this).data('name');let endpoint=$(this).data('endpoint');storage.setItem('active-request',name);storage.setItem('active-endpoint',endpoint);$('.js-active-endpoint').text(endpoint);$('.js-request').removeClass('active');$(this).addClass('active');let requestBody=JSON.parse($(this).find('.js-request-body-template').text());requestBodyEditor.set('loading...');db.requests.where({service:service,name:name}).first().then(function(request){if(request){requestBody=request.body;}}).catch(function(error){console.error(error);}).finally(function(){requestBodyEditor.set(requestBody);});responseEditor.set('loading...');$('.js-response-took').text('...');$('.js-response-time').text('...');$('.js-correlation-id').text('');let cachedResponse={time:'...',receivedAt:'',body:'',correlationId:'',};db.responses.where({service:service,request:name}).first().then(function(response){if(response){cachedResponse=response;}}).catch(function(error){console.error(error);}).finally(function(){responseEditor.set(cachedResponse.body);$('.js-response-took').text(cachedResponse.time);$('.js-response-time').text(cachedResponse.receivedAt?moment(cachedResponse.receivedAt).fromNow():'...');$('.js-correlation-id').text(cachedResponse.correlationId);});});$(document).on('keyup','#request-body-editor',function(){let service=storage.getItem('active-service');let request=storage.getItem('active-request');let body={};try{body=requestBodyEditor.get();}catch(e){return}
try{db.requests.update(request,{body:body}).then(function(updated){if(!updated){db.requests.put({service:service,name:request,body:requestBodyEditor.get()});}});}catch(err){}});$(document).on('keydown','#request-body-editor',function(e){if(e.ctrlKey&&e.keyCode===13){$('.js-send').trigger('click');}});$(document).on('click','.js-reset',function(){let requestName=storage.getItem('active-request');let requestBodyTemplate=$('.js-request[data-name="'+requestName+'"]').find('.js-request-body-template').text();requestBodyEditor.set(JSON.parse(requestBodyTemplate));});$(document).on('click','.js-send',function(){let sendBtn=$('.js-send');sendBtn.prop('disabled',true);let body;try{body=requestBodyEditor.get();}catch(e){alert('Request: '+e.message);sendBtn.prop('disabled',false);}
let contextEditor=$('#request-context-editor textarea');responseEditor.set('loading...');$('.js-response-time').text('...');$('.js-response-took').text('...');$('.js-correlation-id').text('...');let environment=$('.js-environment').val();let service=storage.getItem('active-service');let request=storage.getItem('active-request');let endpoint=storage.getItem('active-endpoint');let bodyText=JSON.stringify(body);let matches=getMatches(bodyText);let promises=[];for(let match in matches){if(!matches.hasOwnProperty(match)){continue;}
if(typeof matches[match].request==='undefined'){continue;}
let source=matches[match];let replace='';promises.push(db.responses.where({service:source.service,request:source.request}).first().then(function(response){if(response){replace=jsonpath.value(response.body,source.path);}}).catch(function(error){console.error(error);}).finally(function(){if(source.type==='int'){match='"'+match+'"';}
bodyText=bodyText.replaceAll(pregQuote(match),replace);}));}
Promise.all(promises).then(function(){body=JSON.parse(bodyText);lines=contextEditor.val().split('\n').filter(function(v){return v!='';});let values={};for(i in lines){if(typeof lines[i]!=='string'){continue;}
pair=lines[i].splitN(':',1);values[pair[0].trim()]=pair[1].trim();}
console.info('Request: '+request+
'\nEnvironment: '+environment+
'\nService: '+service+
'\nEndpoint: '+endpoint+
'\nBody: '+JSON.stringify(body)+
'\nContext: '+JSON.stringify(values));class Response{Body;Time;CorrelationId;}
$.ajax({method:'POST',url:'call',dataType:'json',contentType:'application/json',data:JSON.stringify({environment:environment,service:service,endpoint:endpoint,body:body,context:values,}),success:function(response){let body;try{body=JSON.parse(response.Body);}catch(e){body=response.Body;}
console.info('Response: '+response.Body);$('.js-response-time').text('just now');$('.js-response-took').text(response.Time);$('.js-correlation-id').text(response.CorrelationId);responseEditor.set(body);if(response.Body.length>1000000){body="truncated because of size...";}
db.responses.put({service:service,request:request,receivedAt:moment().format(),time:response.Time,body:body,correlationId:response.CorrelationId,});},error:function(jqXHR,_,errorThrown){alert(errorThrown+': '+jqXHR.responseText);},complete:function(){$('.js-send').prop('disabled',false);}});});});$(document).on('click','.js-favorite',function(){let service=$(this).data('service');if(favorites.indexOf(service)===-1){favorites.push(service);storage.setItem('favorites',JSON.stringify(favorites));let serviceContainer=$(this).closest('.js-service');serviceContainer.parent().prepend(serviceContainer);}
$(this).addClass('js-favorite-on').find('use').attr('href','#star-solid');});$(document).on('click','.js-favorite-on',function(){let service=$(this).data('service');favorites.remove(service);storage.setItem('favorites',JSON.stringify(favorites));$(this).removeClass('js-favorite-on').find('use').attr('href','#star-regular');let serviceContainer=$(this).closest('.js-service');serviceContainer.parent().append(serviceContainer);});$(document).on('click','.js-correlation-id',function(){let env=$('.js-environment').val();let correlationId=$(this).text();copyToClipboard(correlationId);$(this).tooltip('dispose');$(this).tooltip({title:'Copied!'}).tooltip('show');let url=storage.getItem('logs:'+env);if(url!==null){url=url.replaceAll('_correlationId_',correlationId);window.open(url,'_blank');}});$(document).on('mouseleave','.js-correlation-id',function(){$(this).tooltip('dispose');$(this).tooltip({title:'Copy to clipboard'});});$('.js-correlation-id').tooltip({title:'Copy to clipboard'});$(document).on('click','.js-environment-settings',function(){let container=$('.js-environment-logs-urls');container.html('');$('.js-environment option').each(function(){let environment=$(this).val();container.append(Mustache.render(tmplEnvironmentLogsUrl,{environment:environment,url:storage.getItem('logs:'+environment),}));});$('#environmentSettings').modal();});$(document).on('click','.js-environment-settings-save',function(){$('.js-environment-url').each(function(){let environment=$(this).data('environment');let url=$(this).val();if(url===''){storage.removeItem('logs:'+environment);}else{storage.setItem('logs:'+environment,$(this).val());}});});$(document).on({mouseenter:function(){$(this).find('.js-request-duplicate').show();if($(this).find('.js-request').data('static')===true){$(this).find('.js-request-edit').show();$(this).find('.js-request-delete').show();}},mouseleave:function(){$(this).find('.js-request-edit').hide();$(this).find('.js-request-duplicate').hide();$(this).find('.js-request-delete').hide();}},'.js-requests li');$(document).on('click','.js-request-edit',function(e){console.log('edit');e.stopPropagation();});function uniqName(service,name){return name+'.'+1;}
$(document).on('click','.js-request-duplicate',function(e){let request=$(this).closest('.js-request');let name=request.data('name');let service=request.data('service');let endpoint=request.data('endpoint');$('.js-new-request-name').val(uniqName(service,name));$('.js-new-request-name').data('service',service);$('.js-new-request-name').data('endpoint',endpoint);$('.js-new-request-name').data('source-request',name);$('.js-new-request-endpoint').text(endpoint);$('#newRequest').modal();e.stopPropagation();});$(document).on('click','.js-request-delete',function(e){let request=$(this).closest('.js-request');let name=request.data('name');let service=request.data('service');if(request.data('static')!==true){e.stopPropagation();return;}
db.requests.where({service:service,name:name}).delete();db.responses.where({service:service,request:name}).delete();request.remove();e.stopPropagation();});$(document).on('click','.js-new-request-save',function(){let name=$('.js-new-request-name').val();let service=$('.js-new-request-name').data('service');let endpoint=$('.js-new-request-name').data('endpoint');let sourceRequest=$('.js-new-request-name').data('source-request');db.requests.where({service:service,name:name}).first().then(function(exists){if(exists){alert('Request with name "'+name+'" already exists');return;}
let requestBody=JSON.parse('""');db.requests.where({name:sourceRequest}).first().then(function(request){if(request){requestBody=request.body;}}).finally(function(){db.requests.put({name:name,service:service,endpoint:endpoint,body:requestBody,static:true});let container=$('.js-requests[data-service="'+service+'"]');container.find('ul').prepend(Mustache.render(tmplRequest,{name:name,service:service,endpoint:endpoint,requestBodyTemplate:JSON.stringify(requestBody),static:true,}));});}).catch(function(error){console.log(error);});});if(activeEnvironment){$('#environmentSwitcher').val(activeEnvironment).trigger('change');}
if(activeRequest){$('.js-request[data-name="'+activeRequest+'"]').trigger('click');}});</script>`
const JsTemplates = `<script id="tmplService"type="x-tmpl-mustache"><div class="js-service"data-service="{{ name }}"><div class="mt-1 mb-1"><svg class="icon"><use href="#cube"/></svg><svg class="icon clickable js-favorite"data-service="{{ name }}"><use href="#star-regular"/></svg><strong class="collapsed"style="cursor:pointer;"data-toggle="collapse"href="#{{ id }}">{{name}}</strong><span class="badge badge-secondary js-version"></span></div><div class="collapse js-requests"id="{{ id }}"data-service="{{ name }}">...</div></div></script><script id="tmplRequest"type="x-tmpl-mustache"><li><div class="ml-3 pl-1 pr-1 mb-1 js-request"
style="cursor: pointer;display: inline-block;"
data-name="{{ name }}"data-service="{{ service }}"data-endpoint="{{ endpoint }}"data-static="{{ static }}">{{name}}<svg class="js-request-edit icon"style="display:none;"><use href="#edit-solid"/></svg><svg class="js-request-duplicate icon"style="display:none;"><use href="#paste-solid"/></svg><svg class="js-request-delete icon icon-danger"style="display:none;"><use href="#trash-solid"/></svg><pre style="display:none;"class="js-request-body-template">{{requestBodyTemplate}}</pre></div></li></script><script id="tmplEnvironmentLogsUrl"type="x-tmpl-mustache"><div class="form-group"><label for="{{ environment }}Url">{{environment}}</label><textarea class="form-control form-control-sm js-environment-url"data-environment="{{ environment }}"id="{{ environment }}Url"rows="4"spellcheck="false">{{url}}</textarea></div></script>`
