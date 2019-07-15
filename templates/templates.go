package templates

const Index = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"
          integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    <script src="https://kit.fontawesome.com/43cf7afab1.js"></script>
    <title>Hypersomnia</title>
    <style>
        textarea.form-control {
            height: 100%;
        }

        div.active {
            background-color: #444444;
            color: #ffffff;
            -webkit-border-radius: 3px;
            -moz-border-radius: 3px;
            border-radius: 3px;
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

                <textarea class="form-control js-request text-monospace" style="font-size:70%;"></textarea>

        </div>
        <div class="col-sm" style="overflow-x: scroll;overflow-y: scroll;">
            <div>
                <span class="badge badge-secondary js-response-took">...</span>
                <span class="badge badge-secondary js-response-time">...</span>
            </div>
            <div class="js-response text-monospace" style="font-size:70%;white-space: pre-wrap;"></div>
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
<script>
    $(function () {
        let storage = window.localStorage;

        $(document).on('click', '.js-endpoint-toggle', function () {
            let service = $(this).data('service');
            storage.setItem('active-service', service);

            let endpoint = $(this).data('endpoint');
            storage.setItem('active-endpoint', endpoint);
            $('.js-active-endpoint').text(endpoint);
            $('.js-endpoint-toggle').removeClass('active');
            $(this).addClass('active');

            let requestTemplate = storage.getItem(endpoint + ':request');
            if (requestTemplate === null) {
                requestTemplate = $(this).find('.js-endpoint-request-template').text();
            }

            $('.js-response').text('');
            if (cachedResponse = storage.getItem(endpoint + ':response')) {
                $('.js-response').text(cachedResponse);
            }
            $('.js-response-took').text('...');
            if (cachedResponseTook = storage.getItem(endpoint+':responseTook')) {
                $('.js-response-took').text(cachedResponseTook);
            }
            $('.js-response-time').text('...');
            if (cachedResponseTime = storage.getItem(endpoint+':responseTime')) {
                let responseTime = moment(cachedResponseTime);
                $('.js-response-time').text(responseTime.fromNow());
            }

            $('.js-request').val(requestTemplate);
        });

        $(document).on('click', '.js-reset', function () {
            let endpoint = storage.getItem('active-endpoint');
            requestTemplate = $('.js-endpoint-toggle[data-endpoint="'+endpoint+'"]').find('.js-endpoint-request-template').text();
            $('.js-request').val(requestTemplate);
        });

        $(document).on('keyup', '.js-request', function () {
            let endpoint = storage.getItem('active-endpoint');
            storage.setItem(endpoint + ':request', $(this).val());
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

        if (activeEndpoint = storage.getItem('active-endpoint')) {
            $('.js-endpoint-toggle[data-endpoint="' + activeEndpoint + '"]').trigger('click');
        }

        $(document).on('click', '.js-send', function () {
			$('.js-send').prop('disabled', true);
            let service = storage.getItem('active-service');
            let endpoint = storage.getItem('active-endpoint');
            let body;
            try {
                body = JSON.parse($('.js-request').val());
            } catch (e) {
                alert('Request: ' + e.message);
            }

            storage.setItem(endpoint + ':response', '');
            $('.js-response').text('loading...');
            $('.js-response-time').text('...');
            $('.js-response-took').text('...');
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

                    body = JSON.stringify(body, null, 2);
                    let time = response.Time;
					try {
						storage.setItem(endpoint + ':response', body);
					} catch (e) {
						console.log(e.message)
					}
                    storage.setItem(endpoint + ':responseTime', moment().format());
                    storage.setItem(endpoint + ':responseTook', time);

                    $('.js-response-time').text('just now');
                    $('.js-response-took').text(time);
                    $('.js-response').text(body);
                },
				error: function (jqXHR, textStatus, errorThrown) {
					alert("AJAX error");
				},
				complete: function () {
					$('.js-send').prop('disabled', false);
				}
            })
        });

		$(document).delegate('.js-request', 'keydown', function(e) {
			var keyCode = e.keyCode || e.which;
		
			if (keyCode == 9) {
				e.preventDefault();
				var start = this.selectionStart;
				var end = this.selectionEnd;
		
				// set textarea value to: text before caret + tab + text after caret
				$(this).val($(this).val().substring(0, start)
					+ "  "
					+ $(this).val().substring(end));
		
				// put caret at right position again
				this.selectionStart =
					this.selectionEnd = start + 2;
			}
		});
    })
</script>
</body>
</html>
`
