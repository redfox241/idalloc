<?php
/**
 * Thrift RPC - PHPClient
 * @author redfox241
 * @time 2016.10.13
 */

define("ROOT_DIR" , realpath(dirname(__FILE__).'/'));
define("GEN_DIR" , realpath(dirname(__FILE__).'/').'/Protocol');
require_once ROOT_DIR . 'Rpc.php';

$arrUserInfo = new UserClient;
$arrUserInfo->user_id = 2;
$arrUserInfo->user_name = '曹操';
$arrUserInfo->nick_name = '曹孟德';
$arrUserInfo->intro = '宁愿天下人负吾，我不负天下人';

Rpc::call("user","GetUserInfo",$arrUserInfo);

