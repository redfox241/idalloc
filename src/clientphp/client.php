<?php
/**
 * Thrift RPC - PHPClient
 * @author liuxinming
 * @time 2015.5.13
 */

header("Content-type: text/html; charset=utf-8");
$startTime = getMillisecond();//记录开始时间

$ROOT_DIR = realpath(dirname(__FILE__).'/');
$GEN_DIR = realpath(dirname(__FILE__).'/').'/Protocol';
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;
use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\TSocketPool;
use Thrift\Transport\TFramedTransport;
use Thrift\Transport\TBufferedTransport;

$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift',$ROOT_DIR);
$loader->registerDefinition('user\user', $GEN_DIR);
$loader->register();


#require_once ( __DIR__ . "/config.php");

$thriftHost = '127.0.0.1'; //UserServer接口服务器IP
$thriftPort = 9090;            //UserServer端口

$socket = new TSocket($thriftHost,$thriftPort);
$socket->setSendTimeout(10000);#Sets the send timeout.
$socket->setRecvTimeout(20000);#Sets the receive timeout.
//$transport = new TBufferedTransport($socket); #传输方式：这个要和服务器使用的一致 [go提供后端服务,迭代10000次2.6 ~ 3s完成]
$transport = new TFramedTransport($socket); #传输方式：这个要和服务器使用的一致[go提供后端服务,迭代10000次1.9 ~ 2.1s完成，比TBuffer快了点]
$protocol = new TBinaryProtocol($transport);  #传输格式：二进制格式
$client = new \user\user\UserClient($protocol);# 构造客户端

$transport->open();
$socket->setDebug(true);

$s = new \user\user\UserInfo();

$s->user_id = 2;
$s->user_name = '曹操';
$s->nick_name = '曹孟德';
$s->intro = '宁愿天下人负吾，我不负天下人';
$arrResult = $client->CreateNewUser($s);
var_dump( $arrResult );


for($i=1;$i<2;$i++){
    $item = array();

    $item["user_id"] = intval($arrResult);
	$item["user_name"] = "redfox241";
	$item["nick_name"] = "Alex";
	$item["intro"] = "like bird ....";

    $result = $client->GetUserInfo(time(),"php client",$item); # 对服务器发起rpc调用
	$arrUserInfo = (array)( $result[0]) ;

	var_dump( $arrUserInfo );
}


$endTime = getMillisecond();

echo "本次调用用时: :".$endTime."-".$startTime."=".($endTime-$startTime)."毫秒<br>\n";

function getMillisecond() {
    list($t1, $t2) = explode(' ', microtime());
    return (float)sprintf('%.0f', (floatval($t1) + floatval($t2)) * 1000);
}

$transport->close();
