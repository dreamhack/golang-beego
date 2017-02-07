<!DOCTYPE html> 
<html > 
<head> 
<link rel="stylesheet" href="/static/lib/plupload/js/jquery.plupload.queue/css/jquery.plupload.queue.css" type="text/css" media="screen" /> 
<script src="/static/lib/flat/js/jquery-2.0.3.min.js" type="text/javascript"></script>
<script src="/static/lib/plupload/js/plupload.full.min.js" type="text/javascript" ></script>
<script src="/static/lib/plupload/js/jquery.plupload.queue/jquery.plupload.queue.js" type="text/javascript"></script> 
</head> 
<body style="font: 13px Verdana; background: #eee; color: #333"> 
 <div id="html5_uploader">
<form enctype="multipart/form-data" id="html5_uploader">
</form>
</div>
<script type="text/javascript">
// Convert divs to queue widgets when the DOM is ready
$(function() {
    $("#html5_uploader").pluploadQueue({
        // General settings
        runtimes : 'html5',
        url : "/admin/upload",
        chunk_size : '10000mb',
        unique_names : true,
         
        filters : {
            max_file_size : '10000mb',
            mime_types: [
                {title : "Image files", extensions : "jpg,gif,png"},
                {title : "media files", extensions : "mp3,mp4,avi"}
            ]
        },
 
        // Resize images on clientside if we can
        resize : {width : 320, height : 240, quality : 90}
    });

 
});
</script>
</body> 
</html>