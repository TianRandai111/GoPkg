/*
 * @Descripttion: 
 * @version: 
 * @Author: 步荀仙
 * @Date: 2019-09-17 13:58:44
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-09-17 14:22:40
 */
// function fnWarp(){
//     function Touzi() {
//         alert("亲，请关注我们的新产品")
//     };
//     Touzi();
// }
// fnWarp();
//为了防止压缩时，前方语句没有加分号，为了保证语句正常所加的
;(function () {
    function Touzi() {
        alert('亲，请关注我们的新产品')
    };
    Touzi();
})();

// 封闭函数复杂写法
;~function () {
    function Touzi() {
        alert('亲，请关注我们的新产品~')
    };
    Touzi();
}();

;!function () {
    function Touzi() {
        alert('亲，请关注我们的新产品!')
    };
    Touzi();
}();