#!perl
use strict;
use JSON;

# https://api.slack.com/events/api
# Slack doesn't allow us to scrape this page. bummer.
# Data as of Apr 3, 2017
# Note: need to add Pong and Error types. currently added manually
my $content = do { local $/; <DATA> };

# This is really silly, but it works for now
if ($content !~ m{<table class="table table-bordered full_width">(.+)</table>}sm) {
    die q|Couldn't find <table class="table table-bordered full_width">|;
}

my $events_table = $1;
my @events;

# extract <tr></tr>
for my $row ($events_table =~ m{<tr>(.+?)</tr>}gsm) {
    if ($row !~ m{<a href="/events/([^"]+)"}) {
        next;
    }
    my $name = $1;
    my %data = (name => $name, rtm => JSON::false, events => JSON::false);
    if ($row =~ />RTM</) {
        $data{rtm} = JSON::true;
    }
    if ($row =~ />Events API</) {
        $data{events} = JSON::true;
    }
    push @events, \%data;
}

my @data;
foreach my $data (@events) {
    my $r = $data->{name};
    $r =~ s/[^\w]+/_/g;
    $r =~ s/_(\w)/uc($1)/ge;
    $r = ucfirst($r);
    $r =~ s/Url/URL/g;
    $r .= "Type";
    push @data, {name => $r, "value" => $data->{name}, rtm => $data->{rtm}, events => $data->{events} };
}

print JSON->new->pretty->encode(\@data);

__DATA__




<!DOCTYPE html>
<html lang="en" data-locale="en-US">

<head>
	<script>
window.ts_endpoint_url = "https:\/\/slack.com\/beacon\/timing";

(function(e) {
	var n=Date.now?Date.now():+new Date,r=e.performance||{},t=[],a={},i=function(e,n){for(var r=0,a=t.length,i=[];a>r;r++)t[r][e]==n&&i.push(t[r]);return i},o=function(e,n){for(var r,a=t.length;a--;)r=t[a],r.entryType!=e||void 0!==n&&r.name!=n||t.splice(a,1)};r.now||(r.now=r.webkitNow||r.mozNow||r.msNow||function(){return(Date.now?Date.now():+new Date)-n}),r.mark||(r.mark=r.webkitMark||function(e){var n={name:e,entryType:"mark",startTime:r.now(),duration:0};t.push(n),a[e]=n}),r.measure||(r.measure=r.webkitMeasure||function(e,n,r){n=a[n].startTime,r=a[r].startTime,t.push({name:e,entryType:"measure",startTime:n,duration:r-n})}),r.getEntriesByType||(r.getEntriesByType=r.webkitGetEntriesByType||function(e){return i("entryType",e)}),r.getEntriesByName||(r.getEntriesByName=r.webkitGetEntriesByName||function(e){return i("name",e)}),r.clearMarks||(r.clearMarks=r.webkitClearMarks||function(e){o("mark",e)}),r.clearMeasures||(r.clearMeasures=r.webkitClearMeasures||function(e){o("measure",e)}),e.performance=r,"function"==typeof define&&(define.amd||define.ajs)&&define("performance",[],function(){return r}) // eslint-disable-line
})(window);

</script>
<script>;(function() {
'use strict';


window.TSMark = function(mark_label) {
	if (!window.performance || !window.performance.mark) return;
	performance.mark(mark_label);
};
window.TSMark('start_load');


window.TSMeasureAndBeacon = function(measure_label, start_mark_label) {
	if (start_mark_label === 'start_nav' && window.performance && window.performance.timing) {
		window.TSBeacon(measure_label, (new Date()).getTime() - performance.timing.navigationStart);
		return;
	}
	if (!window.performance || !window.performance.mark || !window.performance.measure) return;
	performance.mark(start_mark_label + '_end');
	try {
		performance.measure(measure_label, start_mark_label, start_mark_label + '_end');
		window.TSBeacon(measure_label, performance.getEntriesByName(measure_label)[0].duration);
	} catch (e) {
		
	}
};


if ('sendBeacon' in navigator) {
	window.TSBeacon = function(label, value) {
		var endpoint_url = window.ts_endpoint_url || 'https://slack.com/beacon/timing';
		navigator.sendBeacon(endpoint_url + '?data=' + encodeURIComponent(label + ':' + value), '');
	};
} else {
	window.TSBeacon = function(label, value) {
		var endpoint_url = window.ts_endpoint_url || 'https://slack.com/beacon/timing';
		(new Image()).src = endpoint_url + '?data=' + encodeURIComponent(label + ':' + value);
	};
}
})();
</script>
 

<script>
window.TSMark('step_load');
</script>	<noscript><meta http-equiv="refresh" content="0; URL=/events?nojsmode=1" /></noscript>
<script type="text/javascript">
if(self!==top)window.document.write("\u003Cstyle>body * {display:none !important;}\u003C\/style>\u003Ca href=\"#\" onclick="+
"\"top.location.href=window.location.href\" style=\"display:block !important;padding:10px\">Go to Slack.com\u003C\/a>");

(function() {
	var timer;
	if (self !== top) {
		timer = window.setInterval(function() {
			if (window.$) {
				try {
					$('#page').remove();
					$('#client-ui').remove();
					window.TS = null;
					window.clearInterval(timer);
				} catch(e) {}
			}
		}, 200);
	}
}());

</script>

<script>(function() {
        'use strict';

        window.callSlackAPIUnauthed = function(method, args, callback) {
                var timestamp = Date.now() / 1000;  
                var version = (window.TS && TS.boot_data) ? TS.boot_data.version_uid.substring(0, 8) : 'noversion';
                var url = '/api/' + method + '?_x_id=' + version + '-' + timestamp;
                var req = new XMLHttpRequest();

                req.onreadystatechange = function() {
                        if (req.readyState == 4) {
                                req.onreadystatechange = null;
                                var obj;

                                if (req.status == 200 || req.status == 429) {
                                        try {
                                                obj = JSON.parse(req.responseText);
                                        } catch (err) {
                                                TS.console.warn(8675309, 'unable to do anything with api rsp');
                                        }
                                }

                                obj = obj || {
                                        ok: false,
                                };

                                callback(obj.ok, obj, args);
                        }
                };

                var async = true;
                req.open('POST', url, async);

                var form_data = new FormData();
                var has_data = false;
                Object.keys(args).forEach(function(k) {
                        if (k[0] === '_') return;
                        form_data.append(k, args[k]);
                        has_data = true;
                });

                if (has_data) {
                        req.send(form_data);
                } else {
                        req.send();
                }
        };
})();
</script>
	    <title>API Events | Slack</title>
    <meta name="author" content="Slack">

					
															
		
					
		
		<!-- output_css "core" -->
    <link href="https://a.slack-edge.com/64b3/style/libs/jquery.monkeyScroll.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/87b1/style/libs/ladda-themeless.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/114c/style/libs/bootstrap_plastic.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/54427/style/slack_iconfont.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/dfb74/style/plastic_helpers.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/87c57/style/plastic_typography.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/32c90/style/plastic_layout.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/01035/style/plastic_grid.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/de779/style/plastic_buttons.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/3ea1/style/plastic_forms.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/7f58/style/plastic_modal.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/9e5e/style/plastic_fs_modal.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/dc0ef/style/plastic_menu.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/56ff2/style/plastic_sidebar_menu.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/93cee/style/plastic_typeahead.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/d7ee/style/lazy_filter_select.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/8aa25/style/api.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">

	<!-- output_css "api_docs" -->

	<!-- output_css "regular" -->
    <link href="https://a.slack-edge.com/e293a/style/footer.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">
    <link href="https://a.slack-edge.com/b003/style/libs/lato-2-compressed.css" rel="stylesheet" type="text/css" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)">

	

					
	<meta name="HandheldFriendly" content="true"/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />

					

<meta property="og:type" 	content="website"/>
<meta property="og:site_name" 	content="api.slack.com"/>
<meta property="og:image" 	content="https://a.slack-edge.com/f30f/img/services/api_200.png"/>
<meta property="og:title" 	content="Event Types"/>
<meta property="og:description" content="Discover which event types are distributed in Slack&#039;s RTM &amp; Events APIs"/>
		<meta name="twitter:site" content="@slackapi" />
		<meta name="twitter:card" content="summary" />
		

	<!--[if lt IE 9]>
	<script src="https://a.slack-edge.com/ef0d/js/libs/html5shiv.js"></script>
	<![endif]-->

	
<link id="favicon" rel="shortcut icon" href="https://a.slack-edge.com/66f9/img/icons/favicon-32.png" sizes="16x16 32x32 48x48" type="image/png" />

<link rel="icon" href="https://a.slack-edge.com/0180/img/icons/app-256.png" sizes="256x256" type="image/png" />

<link rel="apple-touch-icon-precomposed" sizes="152x152" href="https://a.slack-edge.com/66f9/img/icons/ios-152.png" />
<link rel="apple-touch-icon-precomposed" sizes="144x144" href="https://a.slack-edge.com/66f9/img/icons/ios-144.png" />
<link rel="apple-touch-icon-precomposed" sizes="120x120" href="https://a.slack-edge.com/66f9/img/icons/ios-120.png" />
<link rel="apple-touch-icon-precomposed" sizes="114x114" href="https://a.slack-edge.com/66f9/img/icons/ios-114.png" />
<link rel="apple-touch-icon-precomposed" sizes="72x72" href="https://a.slack-edge.com/0180/img/icons/ios-72.png" />
<link rel="apple-touch-icon-precomposed" href="https://a.slack-edge.com/66f9/img/icons/ios-57.png" />

<meta name="msapplication-TileColor" content="#FFFFFF" />
<meta name="msapplication-TileImage" content="https://a.slack-edge.com/66f9/img/icons/app-144.png" />	
	
		<script>
			if (window.location.host == 'slack.com' && window.location.search.indexOf('story') < 0) {
				document.cookie = '__cvo_skip_doc=' + escape(document.URL) + '|' + escape(document.referrer) + ';path=/';
			}
		</script>
	

		<script type="text/javascript">
		
		try {
			if(window.location.hash && !window.location.hash.match(/^(#?[a-zA-Z0-9_]*)$/)) {
				window.location.hash = '';
			}
		} catch(e) {}
		
	</script>

	<script type="text/javascript">
						
			window.dataLayer = window.dataLayer || [];
			window.dataLayer.push({
				'gtm.start': Date.now(),
				'event' : 'gtm.js',
				
			});
			var firstScript = document.getElementsByTagName('script')[0];
			var thisScript = document.createElement('script');
			thisScript.async = true;
			thisScript.src = '//www.googletagmanager.com/gtm.js?id=GTM-KH2LPK';
			firstScript.parentNode.insertBefore(thisScript, firstScript);
		
	
				(function(e,c,b,f,d,g,a){e.SlackBeaconObject=d;
		e[d]=e[d]||function(){(e[d].q=e[d].q||[]).push([1*new Date(),arguments])};
		e[d].l=1*new Date();g=c.createElement(b);a=c.getElementsByTagName(b)[0];
		g.async=1;g.src=f;a.parentNode.insertBefore(g,a)
		})(window,document,"script","https://a.slack-edge.com/4e6c/js/libs/slack_beacon.js","sb");
		sb('set', 'token', '3307f436963e02d4f9eb85ce5159744c');

				sb('track', 'pageview');

		
		function track(a) {
			if(window.ga) ga('send','event','web', a);
			if(window.sb) sb('track', a);
		}
		

	</script>


</head>

	<body class="api  feature_related_content">

  	
	
	<header>
		<a id="menu_toggle" class="no_transition show_on_mobile">
			<span class="menu_icon"></span>
			<span class="vert_divider"></span>
		</a>
		<a href="https://api.slack.com/" id="header_logo" class="api"><img alt="Slack API" src="https://a.slack-edge.com/ae57/img/slack_api_logo.png" /></a>
		<div class="header_nav">
						<div class="header_links float_right">
				<a href="/" class='active' data-qa="documentation">Documentation</a>
									<a class="" href="/tutorials" data-qa="tutorials">Tutorials</a>
													<a href="/apps"  data-qa="my_apps">Your Apps</a>
							</div>

		</div>

		<div class="alert_page_stacked">
			
			<div class="alert_page alert_success hidden fade api_alert_page" data-js="api_alert_page">
					<i class="ts_icon ts_icon_check_circle_o"></i>
					Success!			</div>
		</div>

	</header>

	<div id="page" class="layout_two_col">
		
		<div id="page_contents" class="clearfix ">
										

<nav id="api_nav" class="col">

	<div class="sidebar_menu show_on_mobile">
		<h4 class="sidebar_menu_header">My Slack API</h4>
		<ul class="sidebar_menu_list">
			<li><a class="sidebar_menu_list_item is_active" href="/">Documentation</a></li>
			<li><a class="sidebar_menu_list_item " href="/apps">Your Apps</a></li>
		</ul>
	</div>

	<div class="sidebar_menu">
					<h4 class="sidebar_menu_header">Start here</h4>
				<ul class="sidebar_menu_list">
						<li><a class="sidebar_menu_list_item " href="/slack-apps">Building Slack apps</a></li>
						<li><a class="sidebar_menu_list_item " href="/changelog">Recent updates</a></li>
		</ul>
	</div>

			<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">App features</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " href="/internal-integrations">Internal integrations</a></li>
				<li><a class="sidebar_menu_list_item " href="/incoming-webhooks">Incoming webhooks</a></li>
				<li><a class="sidebar_menu_list_item " href="/slash-commands">Slash commands</a></li>
				<li><a class="sidebar_menu_list_item " href="/bot-users">Bot users</a></li>
				<li><a class="sidebar_menu_list_item " href="/enterprise-grid">Enterprise Grid</a></li>
				<li><a class="sidebar_menu_list_item" href="/custom-integrations">Legacy custom integrations</a></li>
			</ul>
		</div>

		<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">Messages</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " href="/docs/messages">Introduction</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/message-formatting">Basic formatting</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/message-attachments">Attaching content</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/message-link-unfurling">Unfurling links</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/message-threading">Threading messages</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/message-buttons">Interactive buttons</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/message-guidelines">Message guidelines</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/messages/builder">Message builder</a></li>
			</ul>
		</div>

		<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">APIs</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " href="/web">Web API</a></li>
				<li><a class="sidebar_menu_list_item " href="/rtm">Real Time Messaging API</a></li>
				<li><a class="sidebar_menu_list_item " href="/events-api">Events API</a></li>
				<li><a class="sidebar_menu_list_item " href="/methods">Methods</a></li>
				<li><a class="sidebar_menu_list_item " href="/types">Object Types</a></li>
				<li><a class="sidebar_menu_list_item is_active" href="/events">Event Types</a></li>
				<li><a class="sidebar_menu_list_item " href="/scim">SCIM API</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/presence">Bot &amp; User Presence</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/deep-linking">Deep linking into clients</a></li>
			</ul>
		</div>

		<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">Slack App Directory</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " href="/docs/slack-apps-checklist">Submission checklist</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/slack-apps-guidelines">Submission guidelines</a></li> 
				<li><a class="sidebar_menu_list_item" href="https://slack.com/apps" target="_blank">App Directory</a></li>
				<li><a class="sidebar_menu_list_item " href="/developer-policies">Developer policies</a></li>
			</ul>
		</div>

		<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">Authentication</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " href="/docs/oauth">Using OAuth 2.0</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/oauth-scopes">Permission Scopes</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/oauth-safety">Safely storing credentials</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/slack-button">Slack Button</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/sign-in-with-slack">Sign in with Slack</a></li>
			</ul>
		</div>

		<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">Keep in touch</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " href="/docs/support">Support and Discussion</a></li>
				<li><a class="sidebar_menu_list_item" href="https://twitter.com/slackapi" target="_blank">@SlackAPI</a></li>
				<li><a class="sidebar_menu_list_item" href="https://medium.com/slack-developer-blog" target="_blank">Platform Blog</a></li>
				<li><a class="sidebar_menu_list_item" href="https://slack.engineering/" target="_blank">Slack Engineering Blog</a></li>
				<li><a class="sidebar_menu_list_item" href="/roadmap">Platform Roadmap</a></li>
				<li><a class="sidebar_menu_list_item" href="/register" target="_blank">Register As a Developer</a></li>
				<li><a class="sidebar_menu_list_item" href="/partner-briefings">Join a Partner Briefing</a></li>
			</ul>
		</div>

		<div class="sidebar_menu">
			<h4 class="sidebar_menu_header">Community</h4>
			<ul class="sidebar_menu_list">
				<li><a class="sidebar_menu_list_item " title="Frequently asked questions" href="/faq">FAQ</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/rate-limits">Rate Limits</a></li>
				<li><a class="sidebar_menu_list_item " href="/community">Libraries and Tools</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/hosting">Hosting Providers</a></li>
				<li><a class="sidebar_menu_list_item " href="/docs/community-code-of-conduct">Code of Conduct</a></li>
				<li><a class="sidebar_menu_list_item " href="/slack-fund">Slack Fund</a></li>
				<li><a class="sidebar_menu_list_item" href="https://trello.com/b/HPpcIqd8/slack-app-ideaboard" target="_blank">Ideaboard</a></li>
				<li><a class="sidebar_menu_list_item" href="https://slack.com/terms-of-service/api">API Terms of Service</a></li>
			</ul>
			<p><a href="/support?ref=api_nav" class="btn">Get Help</a></p>
		</div>
	</nav>			
			<div id="api_main_content" class="col">
<h1>API Event Types</h1>
<div class="card">
	Keep track of everything your app needs to know about by tracking these event types when connecting by websocket to the <a href="/rtm">Real Time Messaging API</a> or by push subscription with the <a href="/events-api">Events API</a>.

	
<table class="table table-bordered full_width">
	<tr>
		<th>Event</th>
		<th>Description</th>
					<th colspan="2">Works with</th>
			</tr>
			<tr>
			<td ><small><a href="/events/accounts_changed" class="bold block">accounts_changed</a></small></td>
			<td >
				<p><small>The list of accounts a user is signed into has changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/bot_added" class="bold block">bot_added</a></small></td>
			<td >
				<p><small>An bot user was added</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/bot_changed" class="bold block">bot_changed</a></small></td>
			<td >
				<p><small>An bot user was changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_archive" class="bold block">channel_archive</a></small></td>
			<td >
				<p><small>A channel was archived</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_created" class="bold block">channel_created</a></small></td>
			<td >
				<p><small>A channel was created</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_deleted" class="bold block">channel_deleted</a></small></td>
			<td >
				<p><small>A channel was deleted</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_history_changed" class="bold block">channel_history_changed</a></small></td>
			<td >
				<p><small>Bulk updates were made to a channel&#039;s history</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_joined" class="bold block">channel_joined</a></small></td>
			<td >
				<p><small>You joined a channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_left" class="bold block">channel_left</a></small></td>
			<td >
				<p><small>You left a channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_marked" class="bold block">channel_marked</a></small></td>
			<td >
				<p><small>Your channel read marker was updated</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_rename" class="bold block">channel_rename</a></small></td>
			<td >
				<p><small>A channel was renamed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/channel_unarchive" class="bold block">channel_unarchive</a></small></td>
			<td >
				<p><small>A channel was unarchived</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/commands_changed" class="bold block">commands_changed</a></small></td>
			<td >
				<p><small>A team slash command has been added or changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/dnd_updated" class="bold block">dnd_updated</a></small></td>
			<td >
				<p><small>Do not Disturb settings changed for the current user</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/dnd_updated_user" class="bold block">dnd_updated_user</a></small></td>
			<td >
				<p><small>Do not Disturb settings changed for a team member</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/email_domain_changed" class="bold block">email_domain_changed</a></small></td>
			<td >
				<p><small>The team email domain has changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/emoji_changed" class="bold block">emoji_changed</a></small></td>
			<td >
				<p><small>A team custom emoji has been added or changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_change" class="bold block">file_change</a></small></td>
			<td >
				<p><small>A file was changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_comment_added" class="bold block">file_comment_added</a></small></td>
			<td >
				<p><small>A file comment was added</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_comment_deleted" class="bold block">file_comment_deleted</a></small></td>
			<td >
				<p><small>A file comment was deleted</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_comment_edited" class="bold block">file_comment_edited</a></small></td>
			<td >
				<p><small>A file comment was edited</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_created" class="bold block">file_created</a></small></td>
			<td >
				<p><small>A file was created</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_deleted" class="bold block">file_deleted</a></small></td>
			<td >
				<p><small>A file was deleted</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_public" class="bold block">file_public</a></small></td>
			<td >
				<p><small>A file was made public</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_shared" class="bold block">file_shared</a></small></td>
			<td >
				<p><small>A file was shared</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/file_unshared" class="bold block">file_unshared</a></small></td>
			<td >
				<p><small>A file was unshared</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/goodbye" class="bold block">goodbye</a></small></td>
			<td >
				<p><small>The server intends to close the connection soon.</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_archive" class="bold block">group_archive</a></small></td>
			<td >
				<p><small>A private channel was archived</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_close" class="bold block">group_close</a></small></td>
			<td >
				<p><small>You closed a private channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_history_changed" class="bold block">group_history_changed</a></small></td>
			<td >
				<p><small>Bulk updates were made to a private channel&#039;s history</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_joined" class="bold block">group_joined</a></small></td>
			<td >
				<p><small>You joined a private channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_left" class="bold block">group_left</a></small></td>
			<td >
				<p><small>You left a private channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_marked" class="bold block">group_marked</a></small></td>
			<td >
				<p><small>A private channel read marker was updated</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_open" class="bold block">group_open</a></small></td>
			<td >
				<p><small>You opened a private channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_rename" class="bold block">group_rename</a></small></td>
			<td >
				<p><small>A private channel was renamed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/group_unarchive" class="bold block">group_unarchive</a></small></td>
			<td >
				<p><small>A private channel was unarchived</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/hello" class="bold block">hello</a></small></td>
			<td >
				<p><small>The client has successfully connected to the server</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/im_close" class="bold block">im_close</a></small></td>
			<td >
				<p><small>You closed a DM</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/im_created" class="bold block">im_created</a></small></td>
			<td >
				<p><small>A DM was created</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/im_history_changed" class="bold block">im_history_changed</a></small></td>
			<td >
				<p><small>Bulk updates were made to a DM&#039;s history</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/im_marked" class="bold block">im_marked</a></small></td>
			<td >
				<p><small>A direct message read marker was updated</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/im_open" class="bold block">im_open</a></small></td>
			<td >
				<p><small>You opened a DM</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/link_shared" class="bold block">link_shared</a></small></td>
			<td >
				<p><small>A message was posted containing one or more links relevant to your application</small></p>
			</td>
			<td ></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/manual_presence_change" class="bold block">manual_presence_change</a></small></td>
			<td >
				<p><small>You manually updated your presence</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/message" class="bold block">message</a></small></td>
			<td >
				<p><small>A message was sent to a channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/message.channels" class="bold block">message.channels</a></small></td>
			<td >
				<p><small>A message was posted to a channel</small></p>
			</td>
			<td ></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/message.groups" class="bold block">message.groups</a></small></td>
			<td >
				<p><small>A message was posted to a private channel</small></p>
			</td>
			<td ></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/message.im" class="bold block">message.im</a></small></td>
			<td >
				<p><small>A message was posted in a direct message channel</small></p>
			</td>
			<td ></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/message.mpim" class="bold block">message.mpim</a></small></td>
			<td >
				<p><small>A message was posted in a multiparty direct message channel</small></p>
			</td>
			<td ></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/pin_added" class="bold block">pin_added</a></small></td>
			<td >
				<p><small>A pin was added to a channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/pin_removed" class="bold block">pin_removed</a></small></td>
			<td >
				<p><small>A pin was removed from a channel</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/pref_change" class="bold block">pref_change</a></small></td>
			<td >
				<p><small>You have updated your preferences</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/presence_change" class="bold block">presence_change</a></small></td>
			<td >
				<p><small>A team member&#039;s presence changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/reaction_added" class="bold block">reaction_added</a></small></td>
			<td >
				<p><small>A team member has added an emoji reaction to an item</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/reaction_removed" class="bold block">reaction_removed</a></small></td>
			<td >
				<p><small>A team member removed an emoji reaction</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/reconnect_url" class="bold block">reconnect_url</a></small></td>
			<td >
				<p><small>Experimental</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/star_added" class="bold block">star_added</a></small></td>
			<td >
				<p><small>A team member has starred an item</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/star_removed" class="bold block">star_removed</a></small></td>
			<td >
				<p><small>A team member removed a star</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/subteam_created" class="bold block">subteam_created</a></small></td>
			<td >
				<p><small>A User Group has been added to the team</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/subteam_self_added" class="bold block">subteam_self_added</a></small></td>
			<td >
				<p><small>You have been added to a User Group</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/subteam_self_removed" class="bold block">subteam_self_removed</a></small></td>
			<td >
				<p><small>You have been removed from a User Group</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/subteam_updated" class="bold block">subteam_updated</a></small></td>
			<td >
				<p><small>An existing User Group has been updated or its members changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_domain_change" class="bold block">team_domain_change</a></small></td>
			<td >
				<p><small>The team domain has changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_join" class="bold block">team_join</a></small></td>
			<td >
				<p><small>A new team member has joined</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_migration_started" class="bold block">team_migration_started</a></small></td>
			<td >
				<p><small>The team is being migrated between servers</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_plan_change" class="bold block">team_plan_change</a></small></td>
			<td >
				<p><small>The team billing plan has changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_pref_change" class="bold block">team_pref_change</a></small></td>
			<td >
				<p><small>A team preference has been updated</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_profile_change" class="bold block">team_profile_change</a></small></td>
			<td >
				<p><small>Team profile fields have been updated</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_profile_delete" class="bold block">team_profile_delete</a></small></td>
			<td >
				<p><small>Team profile fields have been deleted</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_profile_reorder" class="bold block">team_profile_reorder</a></small></td>
			<td >
				<p><small>Team profile fields have been reordered</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
			<tr>
			<td ><small><a href="/events/team_rename" class="bold block">team_rename</a></small></td>
			<td >
				<p><small>The team name has changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/url_verification" class="bold block">url_verification</a></small></td>
			<td >
				<p><small>Verifies ownership of an Events API Request URL</small></p>
			</td>
			<td ></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/user_change" class="bold block">user_change</a></small></td>
			<td >
				<p><small>A team member&#039;s data has changed</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ><span class="btn btn_small pill_btn tools">Events API</span></td>
		</tr>
			<tr>
			<td ><small><a href="/events/user_typing" class="bold block">user_typing</a></small></td>
			<td >
				<p><small>A channel member is typing a message</small></p>
			</td>
			<td ><span class="btn btn_small pill_btn new_feature">RTM</span></td>
			<td ></td>
		</tr>
	</table></div>



	</div>
				</div>
<div id="overlay"></div>
</div>





	


<footer  data-qa="footer">
	<section class="links">
		<div class="grid">
			<div class="col span_1_of_4 nav_col">
				<ul>
					<li class="cat_1">Using Slack</li>
					<li><a href="https://slack.com/is" data-qa="product_footer" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_product">Product</a></li>
					<li><a href="https://slack.com/enterprise" data-qa="enterprise_footer" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_enterprise">Enterprise</a></li>
					<li><a href="https://slack.com/pricing" data-qa="pricing_footer" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_pricing">Pricing</a></li>
					<li><a href="https://get.slack.help/hc" data-qa="support_footer" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_support">Support</a></li>
					<li><a href="https://slack.com/guides" data-qa="getting_started" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_getting_started">Slack Guides</a></li>
					<li><a href="https://slack.com/videoguides" data-qa="video_guides" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_video_guides">Video Guides</a></li>
					<li><a href="https://slack.com/apps" data-qa="app_directory" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_app_directory">App Directory</a></li>
					<li><a href="https://api.slack.com/" data-qa="api" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_api">API</a></li>
				</ul>
			</div>
			<div class="col span_1_of_4 nav_col">
				<ul>
					<li class="cat_2">Slack <ts-icon class="ts_icon_heart"></ts-icon></li>
					<li><a href="https://slack.com/jobs" data-qa="jobs" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_jobs">Jobs</a></li>
					<li><a href="https://slack.com/customers" data-qa="customers" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_customers">Customers</a></li>
					<li><a href="https://slack.com/developers" data-qa="developers" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_developers">Developers</a></li>
					<li><a href="https://slack.com/events" data-qa="events" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_events">Events</a></li>
					<li><a href="http://slackhq.com/" data-qa="blog_footer" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_blog">Blog</a></li>
					<li><a href="https://slack.com/podcast" data-qa="podcast" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_podcast">Podcast</a></li>
					<li><a href="https://slack.shop/" data-qa="slack_shop" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_slack_shop">Slack Shop</a></li>
				</ul>
			</div>
			<div class="col span_1_of_4 nav_col">
				<ul>
					<li class="cat_3">Legal</li>
					<li><a href="https://slack.com/privacy-policy" data-qa="privacy" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_privacy">Privacy</a></li>
					<li><a href="https://slack.com/security" data-qa="security" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_security">Security</a></li>
					<li><a href="https://slack.com/terms-of-service" data-qa="tos" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_tos">Terms of Service</a></li>
					<li><a href="https://slack.com/policies" data-qa="policies" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_policies">Policies</a></li>
				</ul>
			</div>
			<div class="col span_1_of_4 nav_col">
				<ul>
					<li class="cat_4">Handy Links</li>
					<li><a href="https://slack.com/downloads" data-qa="downloads" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_downloads">Download desktop app</a></li>
					<li><a href="https://slack.com/downloads" data-qa="downloads_mobile" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_downloads_mobile">Download mobile app</a></li>
					<li><a href="https://slack.com/brand-guidelines" data-qa="brand_guidelines" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_brand_guidelines">Brand Guidelines</a></li>
					<li><a href="https://slackatwork.com" data-qa="slack_at_work" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_slack_at_work">Slack at Work</a></li>
					<li><a href="https://status.slack.com/" data-qa="status" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_status">Status</a></li>
				</ul>
			</div>
		</div>
	</section>

	<div class="footnote">
		<section>
			<a href="https://slack.com" aria-label="Slack homepage" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_slack_icon"><ts-icon class="ts_icon_slack_pillow" aria-hidden="true"></ts-icon></a>
			<ul>
				<li>
					<a href="https://slack.com/help/contact" data-qa="contact_us" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_contact_us">Contact Us</a>
				</li>
				<li>
					<a href="https://twitter.com/SlackHQ" data-qa="slack_twitter" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_slack_twitter" aria-label="Slack on Twitter"><ts-icon class="ts_icon_twitter" aria-hidden="true"></ts-icon></a>
				</li>
				<li class="yt">
					<a href="https://www.youtube.com/channel/UCY3YECgeBcLCzIrFLP4gblw" data-qa="slack_youtube" data-clog-event="WEBSITE_CLICK" data-clog-params="click_target=footer_slack_youtube" aria-label="Slack on YouTube"><ts-icon class="ts_icon_youtube" aria-hidden="true"></ts-icon></a>
				</li>
			</ul>
		</section>
	</div>
</footer>

<script type="text/javascript">

	/**
	 * A placeholder function that the build script uses to
	 * replace file paths with their CDN versions.
	 *
	 * @param {String} file_path - File path
	 * @returns {String}
	 */
	function vvv(file_path) {

		var vvv_warning = 'You cannot use vvv on dynamic values. Please make sure you only pass in static file paths.';
		if (TS && TS.warn) {
			TS.warn(vvv_warning);
		} else {
			console.warn(vvv_warning);
		}

		return file_path;
	}

	var cdn_url = "https:\/\/slack.global.ssl.fastly.net";
	var inc_js_setup_data = {
		emoji_sheets: {
			apple: 'https://a.slack-edge.com/bfaba/img/emoji_2016_06_08/sheet_apple_64_indexed_256colors.png',
			google: 'https://a.slack-edge.com/f360/img/emoji_2016_06_08/sheet_google_64_indexed_128colors.png',
			twitter: 'https://a.slack-edge.com/bfaba/img/emoji_2016_06_08/sheet_twitter_64_indexed_128colors.png',
			emojione: 'https://a.slack-edge.com/bfaba/img/emoji_2016_06_08/sheet_emojione_64_indexed_128colors.png',
		},
	};
</script><script type="text/javascript">
<!--
	// common boot_data
	var boot_data = {
		start_ms: Date.now(),
		app: '',
		user_id: '',
		no_login: true,
		version_ts: '1491222626',
		version_uid: '71d1a19627d4b0a133c5857bee227e03181e9b9c',
		cache_version: "v16-giraffe",
		cache_ts_version: "v2-bunny",
		redir_domain: 'slack-redir.net',
		signin_url: 'https://slack.com/signin',
		abs_root_url: 'https://slack.com/',
		api_url: '/api/',
		team_url: '',
		image_proxy_url: 'https://slack-imgs.com/',
		beacon_timing_url: "https:\/\/slack.com\/beacon\/timing",
		beacon_error_url: "https:\/\/slack.com\/beacon\/error",
		clog_url: "clog\/track\/",
		api_token: '',
		ls_disabled: false,

		notification_sounds: [{"value":"b2.mp3","label":"Ding","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/b2.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/b2.ogg"},{"value":"animal_stick.mp3","label":"Boing","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/animal_stick.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/animal_stick.ogg"},{"value":"been_tree.mp3","label":"Drop","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/been_tree.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/been_tree.ogg"},{"value":"complete_quest_requirement.mp3","label":"Ta-da","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/complete_quest_requirement.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/complete_quest_requirement.ogg"},{"value":"confirm_delivery.mp3","label":"Plink","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/confirm_delivery.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/confirm_delivery.ogg"},{"value":"flitterbug.mp3","label":"Wow","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/flitterbug.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/flitterbug.ogg"},{"value":"here_you_go_lighter.mp3","label":"Here you go","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/here_you_go_lighter.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/here_you_go_lighter.ogg"},{"value":"hi_flowers_hit.mp3","label":"Hi","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/hi_flowers_hit.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/hi_flowers_hit.ogg"},{"value":"knock_brush.mp3","label":"Knock Brush","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/knock_brush.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/knock_brush.ogg"},{"value":"save_and_checkout.mp3","label":"Whoa!","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/save_and_checkout.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/save_and_checkout.ogg"},{"value":"item_pickup.mp3","label":"Yoink","url":"https:\/\/slack.global.ssl.fastly.net\/7e91\/sounds\/push\/item_pickup.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/item_pickup.ogg"},{"value":"hummus.mp3","label":"Hummus","url":"https:\/\/slack.global.ssl.fastly.net\/7fa9\/sounds\/push\/hummus.mp3","url_ogg":"https:\/\/slack.global.ssl.fastly.net\/46ebb\/sounds\/push\/hummus.ogg"},{"value":"none","label":"None"}],
		alert_sounds: [{"value":"frog.mp3","label":"Frog","url":"https:\/\/slack.global.ssl.fastly.net\/a34a\/sounds\/frog.mp3"}],
		call_sounds: [{"value":"call\/alert_v2.mp3","label":"Alert","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/alert_v2.mp3"},{"value":"call\/incoming_ring_v2.mp3","label":"Incoming ring","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/incoming_ring_v2.mp3"},{"value":"call\/outgoing_ring_v2.mp3","label":"Outgoing ring","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/outgoing_ring_v2.mp3"},{"value":"call\/pop_v2.mp3","label":"Incoming reaction","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/pop_v2.mp3"},{"value":"call\/they_left_call_v2.mp3","label":"They left call","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/they_left_call_v2.mp3"},{"value":"call\/you_left_call_v2.mp3","label":"You left call","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/you_left_call_v2.mp3"},{"value":"call\/they_joined_call_v2.mp3","label":"They joined call","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/they_joined_call_v2.mp3"},{"value":"call\/you_joined_call_v2.mp3","label":"You joined call","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/you_joined_call_v2.mp3"},{"value":"call\/confirmation_v2.mp3","label":"Confirmation","url":"https:\/\/slack.global.ssl.fastly.net\/08f7\/sounds\/call\/confirmation_v2.mp3"}],
		call_sounds_version: "v2",
		max_team_handy_rxns: 5,
		max_channel_handy_rxns: 5,
		max_poll_handy_rxns: 7,
		max_handy_rxns_title_chars: 30,
				default_tz: "America\/Los_Angeles",
				
		feature_tinyspeck: false,
		feature_create_team_google_auth: false,
		feature_flannel_fe: false,
		feature_lazy_load_members_and_bots_everywhere: false,
		feature_flannel_use_canary_sometimes: false,
		feature_deprecate_10_8: false,
		feature_thin_channel_membership: true,
		feature_new_broadcast: false,
		feature_new_threads_view: false,
		feature_message_replies: true,
		feature_page_replies_methods: false,
		feature_message_replies_inline: false,
		feature_threads_paging_flexpane: false,
		feature_threads_slash_cmds: false,
		feature_subteam_members_diff: false,
		feature_a11y_keyboard_shortcuts: false,
		feature_email_ingestion: false,
		feature_msg_consistency: false,
		feature_sli_channel_priority: false,
		feature_thanks: false,
		feature_attachments_inline: false,
		feature_fix_files: true,
		feature_channel_eventlog_client: true,
		feature_no_redirects_in_ssb: true,
		feature_paging_api: false,
		feature_enterprise_frecency: false,
		feature_enterprise_move_channels: true,
		feature_i18n_currencies: false,
		feature_i18n_checkout_updates: false,
		feature_dunning: true,
		feature_basic_analytics: false,
		feature_enterprise_profile_menu_deactivate_2fa: false,
		feature_sso_sandbox: true,
		feature_sso_expose_signed_elements: true,
		feature_js_raf_queue: false,
		feature_channel_alert_words: false,
		feature_shared_channels: false,
		feature_shared_channels_beta: false,
		feature_shared_channels_client: false,
		feature_shared_channels_badges: false,
		feature_shared_channels_invite: false,
		feature_allow_shared_general: false,
		feature_winssb_beta_channel: false,
		feature_presence_sub: true,
		feature_live_support_free_plan: false,
		feature_slackbot_goes_to_college: false,
		feature_newxp_enqueue_message: true,
		feature_lato_2_ssb: true,
		feature_focus_mode: false,
		feature_star_shortcut: false,
		feature_show_jumper_scores: false,
		feature_omit_localstorage_users_bots: false,
		feature_disable_ls_compression: false,
		feature_force_ls_compression: false,
		feature_zendesk_app_submission_improvement: false,
		feature_ignore_code_mentions: false,
		feature_name_tagging_client: false,
		feature_name_tagging_client_extras: false,
		feature_name_tagging_client_search: false,
		feature_use_imgproxy_resizing: false,
		feature_auth_unfurls: true,
		feature_i18n_prod: false,
		feature_share_mention_comment_cleanup: false,
		feature_external_files: false,
		feature_min_web: true,
		feature_electron_memory_logging: false,
		feature_wait_for_all_mentions_in_client: false,
		feature_prev_next_button: false,
		feature_free_inactive_domains: true,
		feature_a11y_tab: false,
		feature_measure_css_usage: false,
		feature_es6_build_pipeline: false,
		feature_take_profile_photo: false,
		feature_arugula: false,
		feature_texty: false,
		feature_texty_takes_over: false,
		feature_texty_browser_substitutions: false,
		feature_texty_copy_paste: false,
		feature_texty_rewrite_on_paste: false,
		feature_texty_search: false,
		feature_toggle_id_translation: false,
		feature_emoji_menu_tuning: false,
		feature_default_shared_channels: false,
		feature_mandatory_shared_channels: false,
		feature_workspace_request: true,
		feature_enable_mdm: false,
		feature_message_menus: false,
		feature_sli_recaps: true,
		feature_sli_recaps_interface: true,
		feature_new_message_click_logging: false,
		feature_user_custom_status: false,
		feature_hide_join_leave: false,
		feature_slim_scrollbar: false,
		feature_flexpane_redesign: true,
		feature_ent_dash_new_modal: false,
		feature_intl_channel_names: true,
		feature_sli_highlight_unreads: false,
		feature_sli_briefing: false,
		feature_deprecate_10_8_modal: true,
		feature_better_app_info: true,
		feature_share_picker: true,
		feature_searcher_dm: true,
		feature_searcher_completions: false,
		feature_platform_disable_rtm: true,
		feature_gsuite_rich_previews: false,
		feature_enterprise_guest_conversion: false,
		feature_scrollback_half_measures: false,
		feature_local_last_event_ts: true,
		feature_client_resize_optimizations: false,
		feature_enterprise_loading_state: false,
		feature_api_admin_page: false,
		feature_delay_channel_history_fetch: false,
		feature_app_permissions_api_site: false,
		feature_app_index: false,
		feature_deprecate_archives_for_admin_channels: false,

	client_logs: {"0":{"numbers":[0],"whitelisted":false},"@scott":{"numbers":[2,4,37,58,67,141,142,389,390,481,488,529,667,773,888,999],"owner":"@scott"},"@eric":{"numbers":[2,23,47,48,65,66,72,73,82,91,93,96,222,365,438,528,552,777,794],"owner":"@eric"},"2":{"owner":"@scott \/ @eric","numbers":[2],"whitelisted":false},"4":{"owner":"@scott","numbers":[4],"whitelisted":false},"5":{"channels":"#dhtml","numbers":[5],"whitelisted":false},"23":{"owner":"@eric","numbers":[23],"whitelisted":false},"sounds":{"owner":"@scott","name":"sounds","numbers":[37]},"37":{"owner":"@scott","name":"sounds","numbers":[37],"whitelisted":true},"47":{"owner":"@eric","numbers":[47],"whitelisted":false},"48":{"owner":"@eric","numbers":[48],"whitelisted":false},"Message History":{"owner":"@scott","name":"Message History","numbers":[58]},"58":{"owner":"@scott","name":"Message History","numbers":[58],"whitelisted":true},"65":{"owner":"@eric","numbers":[65],"whitelisted":false},"66":{"owner":"@eric","numbers":[66],"whitelisted":false},"67":{"owner":"@scott","numbers":[67],"whitelisted":false},"72":{"owner":"@eric","numbers":[72],"whitelisted":false},"73":{"owner":"@eric","numbers":[73],"whitelisted":false},"82":{"owner":"@eric","numbers":[82],"whitelisted":false},"@shinypb":{"owner":"@shinypb","numbers":[88,1000,1989,1990,1996]},"88":{"owner":"@shinypb","numbers":[88],"whitelisted":false},"91":{"owner":"@eric","numbers":[91],"whitelisted":false},"93":{"owner":"@eric","numbers":[93],"whitelisted":false},"96":{"owner":"@eric","numbers":[96],"whitelisted":false},"@steveb":{"owner":"@steveb","numbers":[99]},"99":{"owner":"@steveb","numbers":[99],"whitelisted":false},"Channel Marking (MS)":{"owner":"@scott","name":"Channel Marking (MS)","numbers":[141]},"141":{"owner":"@scott","name":"Channel Marking (MS)","numbers":[141],"whitelisted":true},"Channel Marking (Client)":{"owner":"@scott","name":"Channel Marking (Client)","numbers":[142]},"142":{"owner":"@scott","name":"Channel Marking (Client)","numbers":[142],"whitelisted":true},"222":{"owner":"@eric","numbers":[222],"whitelisted":false},"365":{"owner":"@eric","numbers":[365],"whitelisted":false},"389":{"owner":"@scott","numbers":[389],"whitelisted":false},"390":{"owner":"@scott","numbers":[390],"whitelisted":false},"438":{"owner":"@eric","numbers":[438],"whitelisted":false},"@rowan":{"numbers":[444,666],"owner":"@rowan"},"444":{"owner":"@rowan","numbers":[444],"whitelisted":false},"481":{"owner":"@scott","numbers":[481],"whitelisted":false},"488":{"owner":"@scott","numbers":[488],"whitelisted":false},"528":{"owner":"@eric","numbers":[528],"whitelisted":false},"529":{"owner":"@scott","numbers":[529],"whitelisted":false},"552":{"owner":"@eric","numbers":[552],"whitelisted":false},"dashboard":{"owner":"@ac \/ @bruce \/ @kylestetz \/ @nic \/ @rowan","channels":"#devel-enterprise-fe, #feat-enterprise-dash","name":"dashboard","numbers":[666]},"@ac":{"channels":"#devel-enterprise-fe, #feat-enterprise-dash","name":"dashboard","numbers":[666],"owner":"@ac"},"@bruce":{"channels":"#devel-enterprise-fe, #feat-enterprise-dash","name":"dashboard","numbers":[666],"owner":"@bruce"},"@kylestetz":{"channels":"#devel-enterprise-fe, #feat-enterprise-dash","name":"dashboard","numbers":[666],"owner":"@kylestetz"},"@nic":{"channels":"#devel-enterprise-fe, #feat-enterprise-dash","name":"dashboard","numbers":[666],"owner":"@nic"},"666":{"owner":"@ac \/ @bruce \/ @kylestetz \/ @nic \/ @rowan","channels":"#devel-enterprise-fe, #feat-enterprise-dash","name":"dashboard","numbers":[666],"whitelisted":false},"667":{"owner":"@scott","numbers":[667],"whitelisted":false},"773":{"owner":"@scott","numbers":[773],"whitelisted":false},"777":{"owner":"@eric","numbers":[777],"whitelisted":false},"794":{"owner":"@eric","numbers":[794],"whitelisted":false},"Message Pane Scrolling":{"owner":"@scott","name":"Message Pane Scrolling","numbers":[888]},"888":{"owner":"@scott","name":"Message Pane Scrolling","numbers":[888],"whitelisted":true},"999":{"owner":"@scott","numbers":[999],"whitelisted":false},"1000":{"owner":"@shinypb","numbers":[1000],"whitelisted":false},"Members":{"owner":"@fearon","name":"Members","numbers":[1975]},"@fearon":{"owner":"@fearon","name":"Members","numbers":[1975,98765]},"1975":{"owner":"@fearon","name":"Members","numbers":[1975],"whitelisted":true},"lazy loading":{"owner":"@shinypb","channels":"#devel-flannel","name":"lazy loading","numbers":[1989]},"1989":{"owner":"@shinypb","channels":"#devel-flannel","name":"lazy loading","numbers":[1989],"whitelisted":true},"thin_channel_membership":{"owner":"@shinypb","features":["feature_thin_channel_membership"],"channels":"#devel-infrastructure","name":"thin_channel_membership","numbers":[1990]},"1990":{"owner":"@shinypb","features":["feature_thin_channel_membership"],"channels":"#devel-infrastructure","name":"thin_channel_membership","numbers":[1990],"whitelisted":true},"ms":{"owner":"@shinypb","name":"ms","numbers":[1996]},"1996":{"owner":"@shinypb","name":"ms","numbers":[1996],"whitelisted":true},"@patrick":{"owner":"@patrick","channels":"#dhtml","numbers":[2001,2002,2003,2004]},"2001":{"owner":"@patrick","channels":"#dhtml","numbers":[2001],"whitelisted":false},"dnd":{"owner":"@patrick","channels":"dhtml","name":"dnd","numbers":[2002]},"2002":{"owner":"@patrick","channels":"dhtml","name":"dnd","numbers":[2002],"whitelisted":true},"2003":{"owner":"@patrick","channels":"dhtml","numbers":[2003],"whitelisted":false},"Threads":{"owner":"@patrick","channels":"#feat-threads, #devel-threads","name":"Threads","numbers":[2004]},"2004":{"owner":"@patrick","channels":"#feat-threads, #devel-threads","name":"Threads","numbers":[2004],"whitelisted":true},"mc_sibs":{"name":"mc_sibs","numbers":[9999]},"9999":{"name":"mc_sibs","numbers":[9999],"whitelisted":false},"98765":{"owner":"@fearon","numbers":[98765],"whitelisted":false},"@ainjii":{"owner":"@ainjii","numbers":[8675309]},"8675309":{"owner":"@ainjii","numbers":[8675309],"whitelisted":false}},


		img: {
			app_icon: 'https://a.slack-edge.com/272a/img/slack_growl_icon.png'
		},
		page_needs_custom_emoji: false,
		page_needs_team_profile_fields: false,
		page_needs_enterprise: false	};

	
	
	
	
	
	
	// i18n locale map (eg: maps Slack `ja-jp` to ZD `ja`)
			boot_data.slack_to_zd_locale = {"en-us":"en-us","fr-fr":"fr-fr","de-de":"de","es-es":"es","ja-jp":"ja"};
	
	// client boot data
	
		
	
//-->
</script><script type="text/javascript">



var TS_last_log_date = null;
var TSMakeLogDate = function() {
	var date = new Date();

	var y = date.getFullYear();
	var mo = date.getMonth()+1;
	var d = date.getDate();

	var time = {
	  h: date.getHours(),
	  mi: date.getMinutes(),
	  s: date.getSeconds(),
	  ms: date.getMilliseconds()
	};

	Object.keys(time).map(function(moment, index) {
		if (moment == 'ms') {
			if (time[moment] < 10) {
				time[moment] = time[moment]+'00';
			} else if (time[moment] < 100) {
				time[moment] = time[moment]+'0';
			}
		} else if (time[moment] < 10) {
			time[moment] = '0' + time[moment];
		}
	});

	var str = y + '/' + mo + '/' + d + ' ' + time.h + ':' + time.mi + ':' + time.s + '.' + time.ms;
	if (TS_last_log_date) {
		var diff = date-TS_last_log_date;
		//str+= ' ('+diff+'ms)';
	}
	TS_last_log_date = date;
	return str+' ';
}

var parseDeepLinkRequest = function(code) {
	var m = code.match(/"id":"([CDG][A-Z0-9]{8})"/);
	var id = m ? m[1] : null;

	m = code.match(/"team":"(T[A-Z0-9]{8})"/);
	var team = m ? m[1] : null;

	m = code.match(/"message":"([0-9]+\.[0-9]+)"/);
	var message = m ? m[1] : null;

	return { id: id, team: team, message: message };
}

if ('rendererEvalAsync' in window) {
	var origRendererEvalAsync = window.rendererEvalAsync;
	window.rendererEvalAsync = function(blob) {
		try {
			var data = JSON.parse(decodeURIComponent(atob(blob)));
			if (data.code.match(/handleDeepLink/)) {
				var request = parseDeepLinkRequest(data.code);
				if (!request.id || !request.team || !request.message) return;

				request.cmd = 'channel';
				TSSSB.handleDeepLinkWithArgs(JSON.stringify(request));
				return;
			} else {
				origRendererEvalAsync(blob);
			}
		} catch (e) {
		}
	}
}
</script>



<script type="text/javascript">

	var TSSSB = {
		call: function() {
			return false;
		}
	};

</script>
<script>TSSSB.env = (function() {
	'use strict';

	var v = {
		win_ssb_version: null,
		win_ssb_version_minor: null,
		mac_ssb_version: null,
		mac_ssb_version_minor: null,
		mac_ssb_build: null,
		lin_ssb_version: null,
		lin_ssb_version_minor: null,
		desktop_app_version: null,
	};

	var is_win = (navigator.appVersion.indexOf('Windows') !== -1);
	var is_lin = (navigator.appVersion.indexOf('Linux') !== -1);
	var is_mac = !!(navigator.userAgent.match(/(OS X)/g));

	if (navigator.userAgent.match(/(Slack_SSB)/g) || navigator.userAgent.match(/(Slack_WINSSB)/g)) {
		
		var parts = navigator.userAgent.split('/');
		var version_str = parts[parts.length - 1];
		var version_float = parseFloat(version_str);
		var version_parts = version_str.split('.');
		var version_minor = (version_parts.length == 3) ? parseInt(version_parts[2]) : 0;

		if (navigator.userAgent.match(/(AtomShell)/g)) {
			
			if (is_lin) {
				v.lin_ssb_version = version_float;
				v.lin_ssb_version_minor = version_minor;
			} else if (is_win) {
				v.win_ssb_version = version_float;
				v.win_ssb_version_minor = version_minor;
			} else if (is_mac) {
				v.mac_ssb_version = version_float;
				v.mac_ssb_version_minor = version_minor;
			}

			if (version_parts.length >= 3) {
				v.desktop_app_version = {
					major: parseInt(version_parts[0]),
					minor: parseInt(version_parts[1]),
					patch: parseInt(version_parts[2]),
				};
			}
		} else {
			
			v.mac_ssb_version = version_float;
			v.mac_ssb_version_minor = version_minor;

			
			
			var app_ver = window.macgap && macgap.app && macgap.app.buildVersion && macgap.app.buildVersion();
			var matches = String(app_ver).match(/(?:\()(.*)(?:\))/);
			v.mac_ssb_build = (matches && matches.length == 2) ? parseInt(matches[1] || 0) : 0;
		}
	}

	return v;
})();
</script>




	<!-- output_js "core" -->
<script type="text/javascript" src="https://a.slack-edge.com/cea2d8/js/rollup-core_required_libs.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>
<script type="text/javascript" src="https://a.slack-edge.com/eb065/js/rollup-core_required_ts.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>

	<!-- output_js "core_web" -->
<script type="text/javascript" src="https://a.slack-edge.com/6a697/js/rollup-core_web.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>

	<!-- output_js "secondary" -->
<script type="text/javascript" src="https://a.slack-edge.com/7736d/js/libs/highlight.pack.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>
<script type="text/javascript" src="https://a.slack-edge.com/319b9/js/TS.api_docs.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>
<script type="text/javascript" src="https://a.slack-edge.com/1eea8/js/rollup-secondary_a_required.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>
<script type="text/javascript" src="https://a.slack-edge.com/8fdc2/js/rollup-secondary_b_required.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>

	<!-- output_js "regular" -->
<script type="text/javascript" src="https://a.slack-edge.com/bc013/js/footer.js" crossorigin="anonymous" onload="window._cdn && _cdn.ok(this, arguments)" onerror="window._cdn && _cdn.failed(this, arguments)"></script>


		<script type="text/javascript">
			</script>

	


<script>
	window.boot_data = window.boot_data || {};
	boot_data.no_login = true;
	boot_data.app = 'api';
	boot_data.api_url = '/api/';
	boot_data.app_id = "";
	boot_data.app_name = "";
	boot_data.num_teams = 0;
	boot_data.reserved_commands = {};
	boot_data.num_commands = 0;
	boot_data.sudo_url = "";
	boot_data.api_token = "";

	
	$(function(){
		TS.boot(boot_data);
	});
	
</script>

<script>
	hljs.configure({languages: ['json', 'bash', 'javascript', 'http']});
	hljs.initHighlightingOnLoad();
</script>



<!-- slack-www-hhvm-02059668848c16c19 / 2017-04-03 06:13:41 / v71d1a19627d4b0a133c5857bee227e03181e9b9c / B:H -->

</body>
</html>