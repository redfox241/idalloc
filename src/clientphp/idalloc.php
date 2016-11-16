<?php

require_once "Rpc.php";


$arrInputInfo = [
    "type_name" => 'msg_id'
];

$arrRes = Pgy_Rpc_Rpc::call("idalloc","GenId",$arrInputInfo);

var_dump( $arrRes );

