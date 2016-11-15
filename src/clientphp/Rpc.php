<?php
/**
 * Thrift RPC - PHPClient
 * @author redfox241
 * @time 2016.10.13
 */


define("ROOT_DIR", realpath(dirname(__FILE__) . '/'));
define("GEN_DIR", realpath(dirname(__FILE__) . '/') . '/Protocol');
require_once ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;
use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\TSocketPool;
use Thrift\Transport\TFramedTransport;
use Thrift\Transport\TBufferedTransport;


Class Pgy_Rpc_Rpc
{
    private static $_rpcConf    = false;
    const DEFAULT_DB_CONF_FILE = 'rpc/rpc';

    static private function _init()
    {

    }

    static public function call($servicename, $method, $paramslist)
    {
        $startTime = self::_getMillisecond();//记录开始时间

        $loader = new ThriftClassLoader();
        $loader->registerNamespace('Thrift', ROOT_DIR);
        $loader->registerDefinition($servicename . '\\' . $servicename, GEN_DIR);
        $loader->register();


        if (self::$_rpcConf == false) {
            self::$_rpcConf = Pgy_Common_Conf::getConf(self::DEFAULT_DB_CONF_FILE);
        }

        Bingo_Timer::start($servicename."_".$method);

        $strService = ucfirst($servicename);
        $arrServiceList = self::$_rpcConf["ServiceList"][$strService];

        $thriftHost = trim($arrServiceList["Host"]); //UserServer接口服务器IP
        $thriftPort = trim($arrServiceList["Port"]); //UserServer端口
        $thriftWriteTimeout = intval( $arrServiceList["WriteTimeout"] );
        $thriftReadTimeout = intval( $arrServiceList["ReadTimeout"] );

        if ( empty($thriftHost) || empty($thriftPort) || $thriftReadTimeout == 0 || $thriftWriteTimeout == 0 ){
            Bingo_Log::warning("failed to get rpc_conf,rpc_conf:".serialize($arrServiceList[ $strService ]));
        }

        $socket = new TSocket($thriftHost, $thriftPort);
        $socket->setSendTimeout($thriftWriteTimeout);#Sets the send timeout.
        $socket->setRecvTimeout($thriftReadTimeout);#Sets the receive timeout.
        //$transport = new TBufferedTransport($socket); #传输方式：这个要和服务器使用的一致 [go提供后端服务,迭代10000次2.6 ~ 3s完成]
        $transport = new TFramedTransport($socket); #传输方式：这个要和服务器使用的一致[go提供后端服务,迭代10000次1.9 ~ 2.1s完成，比TBuffer快了点]
        $protocol = new TBinaryProtocol($transport);  #传输格式：二进制格式
        $ClassClient = "\\" . $servicename . "\\" . $servicename . "\\" . ucwords($servicename) . "Client";

        $client = new $ClassClient($protocol);

        $transport->open();
        $socket->setDebug(true);

        $arrResult = $client->$method($paramslist);

        //result object to array
        if (is_array($arrResult) && !empty($arrResult)) {

            foreach ($arrResult as &$item) {
                if (is_object($item)) {
                    $item = (array)$item;
                }
            }
        }

        $endTime = self::_getMillisecond();

        Bingo_Timer::end($servicename."_".$method);
        Bingo_Log::warning( "本次调用用时: :" . $endTime . "-" . $startTime . "=" . ($endTime - $startTime) . "毫秒");

        $transport->close();

        return $arrResult;
    }

    static private function _getMillisecond()
    {
        list($t1, $t2) = explode(' ', microtime());
        return (float)sprintf('%.0f', (floatval($t1) + floatval($t2)) * 1000);
    }

}
