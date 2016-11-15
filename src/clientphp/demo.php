<?php

require_once "Rpc.php";


$arrNewUserInfo = [
    "user_id" => "0",
    "user_name" => "张飞",
    "nick_name" => "张翼德",
    "intro" => "燕人张飞在此",
];

$arrRes = Rpc::call("user","CreateNewUser",$arrNewUserInfo);

var_dump( $arrRes );

$arrUserInfo["user_id"] = $arrRes;

#$arrRes = Rpc::call("user","GetUserInfo",$arrUserInfo);
$arrRes = Rpc::call("user","GetUserInfoByUserId",$arrUserInfo);

var_dump( $arrRes );

$arrUserInfo["last_id"] = "537";

$arrRes = Rpc::call("user","GetUserList",$arrUserInfo);

var_dump( $arrRes );

