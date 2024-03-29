<?php if (!defined('THINK_PATH')) exit(); /*a:4:{s:69:"/www/wwwroot/live/public/../application/admin/view/bdfanyi/index.html";i:1684603915;s:60:"/www/wwwroot/live/application/admin/view/layout/default.html";i:1671020443;s:57:"/www/wwwroot/live/application/admin/view/common/meta.html";i:1671020443;s:59:"/www/wwwroot/live/application/admin/view/common/script.html";i:1671020443;}*/ ?>
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
<title><?php echo (isset($title) && ($title !== '')?$title:''); ?></title>
<meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=no">
<meta name="renderer" content="webkit">
<meta name="referrer" content="never">
<meta name="robots" content="noindex, nofollow">

<link rel="shortcut icon" href="/assets/img/favicon.ico" />
<!-- Loading Bootstrap -->
<link href="/assets/css/backend<?php echo \think\Config::get('app_debug')?'':'.min'; ?>.css?v=<?php echo \think\Config::get('site.version'); ?>" rel="stylesheet">

<?php if(\think\Config::get('fastadmin.adminskin')): ?>
<link href="/assets/css/skins/<?php echo \think\Config::get('fastadmin.adminskin'); ?>.css?v=<?php echo \think\Config::get('site.version'); ?>" rel="stylesheet">
<?php endif; ?>

<!-- HTML5 shim, for IE6-8 support of HTML5 elements. All other JS at the end of file. -->
<!--[if lt IE 9]>
  <script src="/assets/js/html5shiv.js"></script>
  <script src="/assets/js/respond.min.js"></script>
<![endif]-->
<script type="text/javascript">
    var require = {
        config:  <?php echo json_encode($config); ?>
    };
</script>

    </head>

    <body class="inside-header inside-aside <?php echo defined('IS_DIALOG') && IS_DIALOG ? 'is-dialog' : ''; ?>">
        <div id="main" role="main">
            <div class="tab-content tab-addtabs">
                <div id="content">
                    <div class="row">
                        <div class="col-xs-12 col-sm-12 col-md-12 col-lg-12">
                            <section class="content-header hide">
                                <h1>
                                    <?php echo __('Dashboard'); ?>
                                    <small><?php echo __('Control panel'); ?></small>
                                </h1>
                            </section>
                            <?php if(!IS_DIALOG && !\think\Config::get('fastadmin.multiplenav') && \think\Config::get('fastadmin.breadcrumb')): ?>
                            <!-- RIBBON -->
                            <div id="ribbon">
                                <ol class="breadcrumb pull-left">
                                    <?php if($auth->check('dashboard')): ?>
                                    <li><a href="dashboard" class="addtabsit"><i class="fa fa-dashboard"></i> <?php echo __('Dashboard'); ?></a></li>
                                    <?php endif; ?>
                                </ol>
                                <ol class="breadcrumb pull-right">
                                    <?php foreach($breadcrumb as $vo): ?>
                                    <li><a href="javascript:;" data-url="<?php echo $vo['url']; ?>"><?php echo $vo['title']; ?></a></li>
                                    <?php endforeach; ?>
                                </ol>
                            </div>
                            <!-- END RIBBON -->
                            <?php endif; ?>
                            <div class="content">
                                <style type="text/css">
.row textarea {
    margin: 10px 0;
}
</style>
<div>
    <div class="row">
    	<form class="form-horizontal form-ajax" role="form" data-toggle="validator" method="POST" id="form"  action="<?php echo addon_url('bdfanyi/index/index');; ?>">
        <div class="col-xs-9">
            <div class="form-inline" data-toggle="cxselect">
            	<label class="control-label">源语言</label>
                <select class="form-control " data-style="btn-primary" data-live-search="true"  name="from"  style="min-width: 120px;">
                	<option value="auto">自动识别</option>                	
                	<option value="zh">中文</option>
                	<option value="en">英语</option>
                	<option value="kor">韩语</option>
                	<option value="jp">日语</option>
                	<option value="fra">法语</option>
                	<option value="de">德语</option>
                	<option value="ru">俄语</option>
                	<option value="th">泰语</option>
                	<option value="spa">西班牙语</option>
                	<option value="pt">葡萄牙语</option>
                	<option value="ara">阿拉伯语</option>
                	<option value="it">意大利语</option>
                	<option value="nl">荷兰语</option>                	
                	<option value="pl">波兰语</option>                	
                	<option value="bul">保加利亚语</option>
                	<option value="est">爱沙尼亚语</option>
                	<option value="dan">丹麦语</option>
                	<option value="fin">芬兰语</option>
                	<option value="rom">罗马尼亚语</option>
                	<option value="slo">斯洛文尼亚语</option>
                	<option value="swe">瑞典语</option>
                	<option value="hu">匈牙利语</option>
                	<option value="vie">越南语</option>
                	<option value="cht">繁体中文</option>
                </select>
                <label class="control-label">目标语言</label>
                <select class=" form-control" name="to"  style="min-width: 120px;">
                	<option value="zh">中文</option>
                	<option value="en">英语</option>
                	<option value="kor">韩语</option>
                	<option value="jp">日语</option>
                	<option value="fra">法语</option>
                	<option value="de">德语</option>
                	<option value="ru">俄语</option>
                	<option value="th">泰语</option>
                	<option value="spa">西班牙语</option>
                	<option value="pt">葡萄牙语</option>
                	<option value="ara">阿拉伯语</option>
                	<option value="it">意大利语</option>
                	<option value="nl">荷兰语</option>                	
                	<option value="pl">波兰语</option>                	
                	<option value="bul">保加利亚语</option>
                	<option value="est">爱沙尼亚语</option>
                	<option value="dan">丹麦语</option>
                	<option value="fin">芬兰语</option>
                	<option value="rom">罗马尼亚语</option>
                	<option value="slo">斯洛文尼亚语</option>
                	<option value="swe">瑞典语</option>
                	<option value="hu">匈牙利语</option>
                	<option value="vie">越南语</option>
                	<option value="cht">繁体中文</option>

                </select>
            </div>
        </div>

        <div class="col-xs-3 text-right">
            <h6><a class="btn btn-success" href="javascript:;" onclick="$('#form').submit();"><i class="fa fa-play"></i> 启动</a></h6>

        </div>
        <div class="col-xs-12">
            <textarea class="form-control" rows="8" name="q" placeholder="为保证翻译质量，请将单次请求长度控制在 6000 bytes以内。（汉字约为2000个）" data-rule="required" ></textarea>
        </div>
    	</form>
    </div>

    <div style="height: 20px;"></div>
	<div class="row">
        <div class="col-xs-9">
            <div class="form-inline" data-toggle="cxselect">
                <label class="control-label">结果</label>
            </div>
        </div>
        <div class="col-xs-3 text-right">
            <h6><a href="javascript:;" class="btn btn-xs btn-primary bt-geturl" title="获取百度接口url" data-url="<?php echo addon_url('bdfanyi/index/getapiurl');; ?>"><i class="fa fa-pencil"></i>获取百度接口url</a></h6>
        </div>
        <div class="col-xs-12">
            <textarea class="form-control" rows="8" id="q" placeholder="翻译结果会在这里展示"></textarea>
        </div>
	</div>
</div>
<!-- 内容 -->



<div class="tab-content">
    <div class="tab-pane active" id="info9">
        <div class="well">
            百度翻译接口---<span style="color: red;">（注意请求频率限制）</span>
            标准版：
			完全免费，不限使用字符量（QPS=1）；
			高级版：
			每月前200万字符免费，超出后仅收取超出部分费用（QPS=10），49元/百万字符；
			尊享版：
			每月前200万字符免费，超出后仅收取超出部分费用（QPS=100），49元/百万字符；

        </div>
        <div class="panel panel-default">
            <div class="panel-heading">
                <strong>
                    请求地址
                </strong>
            </div>
            <div class="panel-body">
                <?php echo addon_url('bdfanyi/index/index',[],true,true);; ?>
            </div>
        </div>
        <div class="panel panel-default">
            <div class="panel-heading">
                <strong>
                    请求方式
                </strong>
            </div>
            <div class="panel-body">
                GET/POST
            </div>
        </div>
        <div class="panel panel-default">
            <div class="panel-heading">
                <strong>
                    参数
                </strong>
                (语种代码参考地址)<a href="https://api.fanyi.baidu.com/product/113" target="_BLANK">https://api.fanyi.baidu.com/product/113</a>
            </div>
            <div class="panel-body">
                <table class="table table-hover">
                    <thead>
                        <tr>
                            <th>
                                名称
                            </th>
                            <th>
                                类型
                            </th>
                            <th>
                                必选
                            </th>
                            <th>
                                描述
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>
                                from
                            </td>
                            <td>
                                string
                            </td>
                            <td>
                                是
                            </td>
                            <td>
                                源语言
                            </td>
                        </tr>
                        <tr>
                            <td>
                                to
                            </td>
                            <td>
                                string
                            </td>
                            <td>
                                是
                            </td>
                            <td>
                                目标语言
                            </td>
                        </tr>
                        <tr>
                            <td>
                                q
                            </td>
                            <td>
                                string
                            </td>
                            <td>
                                是
                            </td>
                            <td>
                                翻译内容
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <div class="panel panel-default">
            <div class="panel-heading">
                <strong>
                    返回示例
                </strong>
            </div>

            <div class="panel-body">

<pre id="response11">{
  <span class="key">"code":</span> <span class="number">1</span>,
  <span class="key">"msg":</span> <span class="string">"成功"</span>,
  <span class="key">"wait":</span> <span class="string">3</span>,
  <span class="key">"data":</span>[{
	<span class="key">"from":</span> <span class="number">"en"</span>,
	<span class="key">"to":</span> <span class="number">"zh"</span>,
	<span class="key">"url":</span> <span class="number">"https://fanyi-api.baidu.com/api/trans/vip/translat....."</span>,
	<span class="key">"trans_result":</span>[{
		<span class="key">"src":</span> <span class="number">"hello"</span>,
		<span class="key">"dst":</span> <span class="number">"你好"</span>,
	}],
  }]
}</pre>

            </div>


        </div>
    </div>
    


</div>












                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <script src="/assets/js/require<?php echo \think\Config::get('app_debug')?'':'.min'; ?>.js" data-main="/assets/js/require-backend<?php echo \think\Config::get('app_debug')?'':'.min'; ?>.js?v=<?php echo htmlentities($site['version']); ?>"></script>
    </body>
</html>
