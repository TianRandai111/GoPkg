/*
 * @Descripttion: 
 * @version: 
 * @Author: 步荀仙
 * @Date: 2019-12-05 17:42:27
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-12-05 19:43:49
 */
  
 $(function(){
    var $name = $('#user_name');

    var $error_name = true;
     
    $name.click(function(){
        $(this).next().hide();
    }).blur(function(){
        fn_check_user_name();
    })

    
    function fn_check_user_name (params) {
        var val =$name.val();
        var reUser = /^\w{6,20}$/
        // 判断输入框是否为空
        if(val==''){
            $error_name = true; 
            $name.next().html('输入框不能为空！').show();
            return;
        }

        if (reUser.test(val)){
            $error_name = false;
            $name.next().hide();
        }else {
            $error_name = true;
            $name.next().html('用户名是6-20位的数字字母或者下划线！').show();
        }
    }
 })