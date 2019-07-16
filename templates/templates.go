package templates

const Index = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
	<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/jsoneditor/6.1.0/jsoneditor.min.css">
    <script src="https://kit.fontawesome.com/43cf7afab1.js"></script>
    <title>Hypersomnia</title>
    <style>
        div.active {
            background-color: #444444;
            color: #ffffff;
            -webkit-border-radius: 3px;
            -moz-border-radius: 3px;
            border-radius: 3px;
        }

		.jsoneditor {
			line-height: auto;
		}

		.jsoneditor-schema-error, div.jsoneditor td, div.jsoneditor textarea, div.jsoneditor th, div.jsoneditor-field, div.jsoneditor-value {
			font-size: 10px;
		}

		.ace-jsoneditor.ace_editor {
			font-size: 10px!important;
		}

		#request-editor, #response-editor {
			height: 100%;
		}
    </style>
</head>
<body>
<div class="container-fluid p-2 pl-4">
    <div class="row">
        <div class="col-sm">
            <div class="mt-1 mb-1">
                Services:
            </div>
            {{range .Services}}
                <div>
                    <div style="cursor:pointer;" class="mt-1 mb-1" data-toggle="collapse" href="#{{id .Name}}"
                         role="button"><strong>
                            <i class="fas fa-cube"></i> {{.Name}}
                        </strong></div>
                    <div class="mb-4 collapse show" id="{{id .Name}}">
                        {{ $service := .Name }}
                        <ul class="list-unstyled">
                            {{range .Endpoints}}
                                <li><div class="ml-3 pl-1 pr-1 mb-2 js-endpoint-toggle" style="cursor: pointer;display: inline-block;"
                                    data-service="{{$service}}" data-endpoint="{{.Name}}">
                                    {{.Name}}
                                    <pre style="display:none;"
                                         class="js-endpoint-request-template">{{- formatEndpoint .Request -}}</pre>
                                </div></li>
                            {{else}}
                                <li class="pl-3 mb-2 text-muted">(no endpoints)</li>
                            {{end}}
                        </ul>
                    </div>
                </div>
            {{else}}
                <div>
                    <strong>no services are running</strong><br>
                    <small>if you think this is wrong, check your registry choice</small>
                </div>
            {{end}}
        </div>
        <div class="col-sm">
            <div class="mb-2 clearfix">
                <div class="js-active-endpoint float-left pt-1 pb-1"></div>
                <button class="btn btn-sm btn-primary js-send float-right">Send</button>
                <button class="btn btn-sm btn-secondary js-reset float-right mr-2">Reset</button>
            </div>
			<div id="request-editor"></div>
        </div>
        <div class="col-sm">
            <div class="mb-2 mt-3 clearfix">
				<span class="badge badge-secondary js-response-time float-right">...</span>                
				<span class="badge badge-secondary js-response-took float-right mr-2">...</span>
            </div>
			<div id="response-editor"></div>
        </div>
    </div>


</div>

<!-- Optional JavaScript -->
<!-- jQuery first, then Popper.js, then Bootstrap JS -->
<script src="https://code.jquery.com/jquery-3.3.1.min.js" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"
        integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1"
        crossorigin="anonymous"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"
        integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"
        crossorigin="anonymous"></script>
<script src="https://momentjs.com/downloads/moment.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jsoneditor/6.1.0/jsoneditor.min.js"></script>
<script src="https://unpkg.com/dexie@latest/dist/dexie.js"></script>
<script src="https://cdn.jsdelivr.net/npm/jsonpath@1.0.2/jsonpath.min.js"></script>
<script>
let storage = window.localStorage;

let db = new Dexie("hypersomnia");
db.version(1).stores({
    requests: 'endpoint,body',
    responses: 'endpoint,time,receivedAt,body'
});

$(function () {
    let requestEditor = new JSONEditor(document.getElementById("request-editor"), {
        enableSort: false,
        enableTransform: false,
		mode: 'code',
    });
    let responseEditor = new JSONEditor(document.getElementById("response-editor"), {
        mode: 'code',
    });

	function setRequestBody(body) {
		
	}

    $(document).on('click', '.js-endpoint-toggle', function () {
        let service = $(this).data('service');
        storage.setItem('active-service', service);

        let endpoint = $(this).data('endpoint');
        storage.setItem('active-endpoint', endpoint);
        $('.js-active-endpoint').text(endpoint);
        $('.js-endpoint-toggle').removeClass('active');
        $(this).addClass('active');

        let requestBody = JSON.parse($(this).find('.js-endpoint-request-template').text());
		requestEditor.set('loading...');
        db.requests.where('endpoint').equals(endpoint).first().then(function(request) {
            if (request) {
				requestBody = request.body;
			}
        }).catch(function(error){
            console.log(error);
        }).finally(function () {
			requestEditor.set(requestBody);
		});

		responseEditor.set('loading...');
		$('.js-response-took').text('...');
		$('.js-response-time').text('...');
		let cachedResponse = {
			time: '...',
			receivedAt: '',
			body: ''
		};
        db.responses.where('endpoint').equals(endpoint).first().then(function(response) {
            if (response) {
				cachedResponse = response;
			}
        }).catch(function(error){
            console.log(error);
        }).finally(function () {
			responseEditor.set(cachedResponse.body);
			$('.js-response-took').text(cachedResponse.time);
            $('.js-response-time').text(moment(cachedResponse.receivedAt).fromNow());
		});
    });

    $(document).on('click', '.js-reset', function () {
        let endpoint = storage.getItem('active-endpoint');
        requestTemplate = $('.js-endpoint-toggle[data-endpoint="'+endpoint+'"]').find('.js-endpoint-request-template').text();
        requestEditor.set(JSON.parse(requestTemplate));
    });

    $(document).on('keyup', '#request-editor', function () {
        let endpoint = storage.getItem('active-endpoint');
        db.requests.put({endpoint: endpoint, body: requestEditor.get()});
    });

    $('.collapse')
        .on('show.bs.collapse', function () {
            let id = $(this).attr('id');
            storage.setItem(id + ':show', true);
        })
        .on('hide.bs.collapse', function () {
            let id = $(this).attr('id');
            storage.setItem(id + ':show', false);
        })
    ;

    $('.collapse').each(function () {
        let id = $(this).attr('id');
        if (storage.getItem(id + ':show') === 'true') {
            $(this).collapse('show');
        } else {
            $(this).collapse('hide');
        }
    });

	function getMatches(string) {
        var matches = [];
        var match;
        let regex = /Response\.(.+?)\((.+?),?(\s?int)?\)/g;
        while (match = regex.exec(string)) {
            matches[match[0]] = {
                endpoint: match[1],
                path: match[2],
				type: match[3] || 'string',
            };
        }
        return matches;
    }

	function pregQuote (str, delimiter) {
		// Quote regular expression characters plus an optional character  
		// 
		// version: 1107.2516
		// discuss at: http://phpjs.org/functions/preg_quote
		// +   original by: booeyOH
		// +   improved by: Ates Goral (http://magnetiq.com)
		// +   improved by: Kevin van Zonneveld (http://kevin.vanzonneveld.net)
		// +   bugfixed by: Onno Marsman
		// +   improved by: Brett Zamir (http://brett-zamir.me)
		// *     example 1: preg_quote("$40");
		// *     returns 1: '\$40'
		// *     example 2: preg_quote("*RRRING* Hello?");
		// *     returns 2: '\*RRRING\* Hello\?'
		// *     example 3: preg_quote("\\.+*?[^]$(){}=!<>|:");
		// *     returns 3: '\\\.\+\*\?\[\^\]\$\(\)\{\}\=\!\<\>\|\:'
		return (str + '').replace(new RegExp('[.\\\\+*?\\[\\^\\]$(){}=!<>|:\\' + (delimiter || '/') + '-]', 'g'), '\\$&');
	}

	String.prototype.replaceAll = function(search, replacement) {
		var target = this;
		return target.replace(new RegExp(search, 'g'), replacement);
	};

    $(document).on('click', '.js-send', function () {
        $('.js-send').prop('disabled', true);
        let service = storage.getItem('active-service');
        let endpoint = storage.getItem('active-endpoint');
        let body;
        try {
            body = requestEditor.get();
        } catch (e) {
            alert('Request: ' + e.message);
        }

        responseEditor.set('loading...')
        $('.js-response-time').text('...');
        $('.js-response-took').text('...');

		bodyText = JSON.stringify(body);
		let matches = getMatches(bodyText);

		let promises = [];
		for (let match in matches) {
			let source = matches[match];
			let replace = '';
			
			promises.push(db.responses.where('endpoint').equals(source.endpoint).first().then(function(response) {
				if (response) {
					replace = jsonpath.value(response.body, source.path);
				}
			}).catch(function(error){
				console.log(error);
			}).finally(function () {
				if (source.type === 'int') {
					match = '"' + match + '"';					
				}
				bodyText = bodyText.replaceAll(pregQuote(match), replace);
			}));
		}
		
		Promise.all(promises).then(function() {
  			body = JSON.parse(bodyText);
			console.log(body);
			$.ajax({
				method: 'POST',
				url: 'call',
				dataType: 'json',
				contentType: 'application/json',
				data: JSON.stringify({
					service: service,
					endpoint: endpoint,
					body: body,
				}),
				success: function (response) {
					let body;
					try {
						body = JSON.parse(response.Body);
					} catch (e) {
						body = response.Body;
					}
	
					$('.js-response-time').text('just now');
					$('.js-response-took').text(response.Time);
					responseEditor.set(body);
	
					db.responses.put({
						endpoint: endpoint,
						receivedAt: moment().format(),
						time: response.Time,
						body: body
					});
				},
				error: function (jqXHR, textStatus, errorThrown) {
					alert("AJAX error");
				},
				complete: function () {
					$('.js-send').prop('disabled', false);
				}
			});
		});
    });

    $('#request-editor').keydown(function (e) {
        if (e.ctrlKey && e.keyCode === 13) {
            $('.js-send').trigger('click');
        }
    });

    let activeEndpoint = storage.getItem('active-endpoint');
    if (activeEndpoint) {
        $('.js-endpoint-toggle[data-endpoint="' + activeEndpoint + '"]').trigger('click');
    }
});
</script>
</body>
</html>
`
