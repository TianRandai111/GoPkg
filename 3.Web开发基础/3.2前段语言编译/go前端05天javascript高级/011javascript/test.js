/*
 * @Descripttion: 
 * @version: 
 * @Author: 步荀仙
 * @Date: 2019-09-17 10:43:49
 * @LastEditors: 步荀仙
 * @LastEditTime: 2019-09-17 11:00:22
 */
window.onload = function() {
     var oBt = this.document.getElementById('btn');
     var iNum01 = 12;

     oBt.onclick = function() {
         var iNum02 = 24;
         var iRs = fnAdd(iNum01,iNum02);
         alert(iRs);
     }
     function fnAdd(a,b) {
         var iRs2 = a + b;
         return iRs2;
     }
}