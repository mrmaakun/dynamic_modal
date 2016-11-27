$('#modalButton').click(function() {
	$.get("modal", function(data, status){
		        $(data).replaceAll(".modal-dialog")
	});
});