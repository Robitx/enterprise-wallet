function getRequest(item, func) {
  var req = new XMLHttpRequest()

  req.onreadystatechange = function() {
    if(req.readyState == 4) {
      func(req.response)
    }
  }
  req.open("GET", "/GET?request=" + item, true)
  req.send()
}

function postRequest(request, jsonObj, func) {
  var req = new XMLHttpRequest()

  req.onreadystatechange = function() {
    if(req.readyState == 4) {
      func(req.response)
    }
  }

  var formData = new FormData();
  formData.append("request", request)
  formData.append("json", jsonObj)

  req.open("POST", "/POST")
  req.send(formData)
}

// Jquery on all pages
$(window).load(function() {
    updateBalances()
});
setInterval(updateBalances,5000);

// Updates total balances on the page
function updateBalances() {
  getRequest("balances", function(resp){
        obj = JSON.parse(resp)
        if (obj.Error != "none") {
          return
        } 

        $("#ec-balance").text(obj.Content.EC)
        fcBal = formatFC(obj.Content.FC)
        $("#factoid-balance").text(fcBal[0] + ".")
        if(fcBal.length > 1) {
          $("#factoid-balance-trailing").text(fcBal[1])
        } else {
          $("#factoid-balance-trailing").text(0)
        }
  })
}

function formatFC(fcBalance){
  dec = (fcBalance/1e8).toFixed(8)
  decStr = dec.toString()
  decSplit = decStr.split(".")

  return decSplit
}