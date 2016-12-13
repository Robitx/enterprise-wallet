var $jscomp={scope:{},checkStringArgs:function(a,c,b){if(null==a)throw new TypeError("The 'this' value for String.prototype."+b+" must not be null or undefined");if(c instanceof RegExp)throw new TypeError("First argument to String.prototype."+b+" must not be a regular expression");return a+""}};
$jscomp.defineProperty="function"==typeof Object.defineProperties?Object.defineProperty:function(a,c,b){if(b.get||b.set)throw new TypeError("ES3 does not support getters and setters.");a!=Array.prototype&&a!=Object.prototype&&(a[c]=b.value)};$jscomp.getGlobal=function(a){return"undefined"!=typeof window&&window===a?a:"undefined"!=typeof global&&null!=global?global:a};$jscomp.global=$jscomp.getGlobal(this);
$jscomp.polyfill=function(a,c,b,d){if(c){b=$jscomp.global;a=a.split(".");for(d=0;d<a.length-1;d++){var e=a[d];e in b||(b[e]={});b=b[e]}a=a[a.length-1];d=b[a];c=c(d);c!=d&&null!=c&&$jscomp.defineProperty(b,a,{configurable:!0,writable:!0,value:c})}};
$jscomp.polyfill("String.prototype.startsWith",function(a){return a?a:function(a,b){var c=$jscomp.checkStringArgs(this,a,"startsWith");a+="";for(var e=c.length,f=a.length,h=Math.max(0,Math.min(b|0,c.length)),g=0;g<f&&h<e;)if(c[h++]!=a[g++])return!1;return g>=f}},"es6-impl","es3");$(window).load(function(){LoadTransactions()});ContentLen=CurrentCount=0;LoopstopIncrement=15;Loopstop=20;var Transactions=[];Done=!1;
function LoadTransactions(){getRequest("related-transactions",function(a){obj=JSON.parse(a);if("none"==obj.Error&&null==obj.Content)setTimeout(function(){LoadTransactions()},5E3);else if($("#loading-container").remove(),"none"!=obj.Error)SetGeneralError(obj.Error);else if(ContentLen=obj.Content.length,Transactions=obj.Content,0<obj.Content.length&&"empty"==obj.Content[0].TxID)SetGeneralError("Your addresses do not have any transactions in the blockchain");else if(0!=obj.Content.length)for(ContentLen<
Loopstop&&(Loopstop=ContentLen);CurrentCount<Loopstop;CurrentCount++)AppendNewTransaction(Transactions[CurrentCount],CurrentCount)})}Empty=!1;
function LoadCached(){if(!Empty)for(ContentLen<2*Loopstop&&(j=JSON.stringify({Current:ContentLen,More:5*Loopstop}),postRequest("more-cached-transaction",j,function(a){obj=JSON.parse(a);"none"==obj.Error&&null!=obj.Content&&(0==obj.Content.length?Empty=!0:(ContentLen+=obj.Content.length,Transactions=Transactions.concat(obj.Content)))})),ContentLen<Loopstop&&(Loopstop=ContentLen);CurrentCount<Loopstop;CurrentCount++)AppendNewTransaction(Transactions[CurrentCount],CurrentCount)}
function AppendNewTransaction(a,c){if(1==a.Action[0]){pic="sent";amt=0;token="FCT";addrs="";for(var b=0;b<a.Inputs.length;b++)""!=a.Inputs[b].Name&&(addrs+='<div class="nick">'+a.Inputs[b].Name+'<pre class="show-for-large"> ('+a.Inputs[b].Address+")</pre></div>",amt+=a.Inputs[b].Amount/1E8);appendTrans(pic,c,-1*amt,token,a.Date,addrs)}if(1==a.Action[1]){pic="received";amt=0;token="FCT";addrs="";for(b=0;b<a.Outputs.length;b++)""!=a.Outputs[b].Name&&a.Outputs[b].Address.startsWith("FA")&&(addrs+='<div class="nick">'+
a.Outputs[b].Name+'<pre class="show-for-large percent95"> ('+a.Outputs[b].Address+")</pre></div>",amt+=a.Outputs[b].Amount/1E8);appendTrans(pic,c,amt,token,a.Date,addrs)}if(1==a.Action[2]){pic="converted";amt=0;token="FCT";addrs="";for(b=0;b<a.Outputs.length;b++)""!=a.Outputs[b].Name&&a.Outputs[b].Address.startsWith("EC")&&(addrs+='<div class="nick">'+a.Outputs[b].Name+'<pre class="show-for-large percent95"> ('+a.Outputs[b].Address+")</pre></div>",amt+=a.Outputs[b].Amount/1E8);appendTrans(pic,c,amt,
token,a.Date,addrs)}}function appendTrans(a,c,b,d,e,f){$("#transaction-list").append('<tr><td><a data-toggle="transDetails"><i class="transIcon '+a+'"><img src="img/transaction_'+a+'.svg" class="svg"></i></a></td><td>'+e+' : <a value="'+c+'" id="transaction-link" data-toggle="transDetails">'+a.capitalize()+"</a>"+f+'</td><td style="word-wrap: break-word;">'+Number(b.toFixed(4))+" "+d+"</td></tr>")}
$("main").bind("scroll",function(){$("main").scrollTop()+$("main").innerHeight()>=.8*$("main").prop("scrollHeight")&&(Loopstop+=LoopstopIncrement,LoadCached())});LOCAL_EXPLORER=!0;
$("#transaction-list").on("click","#transaction-link",function(){port=$("#controlpanel-port").text();setTransDetails($(this).attr("value"));$("#transDetails #link").attr("href","http://explorer.factom.org/tx/"+Transactions[$(this).attr("value")].TxID);-1==port?(LOCAL_EXPLORER=!1,$("#transDetails #local-link").addClass("disabled-input"),$("#transDetails #local-link").prop("disabled",!0),$("#transDetails #local-link").attr("href","")):$("#transDetails #local-link").attr("href","http://localhost:"+port+
"/search?input="+Transactions[$(this).attr("value")].TxID+"&type=facttransaction")});$("#transDetails #local-link").click(function(a){LOCAL_EXPLORER||a.preventDefault()});
function setTransDetails(a){trans=Transactions[a];$("#trans-detail-txid").text(trans.TxID);$("#trans-details-inputs").html("");for(a=0;a<trans.Inputs.length;a++)$("#trans-details-inputs").append("<tr><td>"+trans.Inputs[a].Address+"</td><td>"+FCTNormalize(trans.Inputs[a].Amount)+" FCT</td></tr>");$("#trans-details-outputs").html("");for(a=0;a<trans.Outputs.length;a++)$("#trans-details-outputs").append("<tr><td>"+trans.Outputs[a].Address+"</td><td>"+FCTNormalize(trans.Outputs[a].Amount)+" FCT</td></tr>");
$("#trans-details-outputs").append('<tr class="total"><td> Total </td><td>'+FCTNormalize(trans.TotalInput)+" FCT</td></tr>");$("#total-transacted").text(FCTNormalize(trans.TotalECOutput+trans.TotalFCTOutput));$("#trans-date").text(trans.Date+" at "+trans.Time)}String.prototype.capitalize=function(){return this.charAt(0).toUpperCase()+this.slice(1)};
