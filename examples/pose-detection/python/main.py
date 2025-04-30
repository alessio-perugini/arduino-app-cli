import gradio as gr
from appslab.modules.posedetection import PoseDetection

### Optional gradio styling and theming #
theme = gr.themes.Ocean(
    primary_hue="green",
    secondary_hue="green",
)

js_func = """
function refresh() {
    const url = new URL(window.location);

    if (url.searchParams.get('__theme') !== 'dark') {
        url.searchParams.set('__theme', 'dark');
        window.location.href = url.href;
    }
}
"""
#########################################

pose = PoseDetection(input_image_size=160)

def pose_det_cv(frame):
    global pose

    boxes = []
    objects = []
    cfd = []
    coordinates = []
    detection = []

    out = pose.detect(frame)
    if out and "detection" in out:
        for i, obj_det in enumerate(out["detection"]):
            if "class_name" not in obj_det or "bounding_box_xyxy" not in obj_det or "confidence" not in obj_det:
                continue

            bl = f"block {i}"
            if "keypoints" in obj_det:
                knots = obj_det["keypoints"]
                for k in knots:
                    if "x" in k and "y" in k:
                        boxes.append(bl)
                        objects.append(k["name"])
                        coordinates.append(f"{k['x']:.1f} {k['y']:.1f}")
                        cfd.append(f"{k['confidence']:.1f}%")
                        detection.append(obj_det["class_name"])

    detection_data = []
    for box, det, obj, coord, conf in zip(boxes, detection, objects, coordinates, cfd):
        detection_data.append([box, det, obj, coord, conf])

    return pose.draw_bounding_boxes(frame, out), detection_data

with gr.Blocks(theme=theme,
               title="Arduino vision object detection demo",
               js=js_func,
               delete_cache=(30, 60),
               css="footer{display:none !important}") as demo:
    
    gr.HTML(value="<img src='https://upload.wikimedia.org/wikipedia/commons/8/87/Arduino_Logo.svg' width='100px' style='float:right'>", elem_id="arduino_logo")
    gr.Markdown("# Vision pose detection with Yolo")

    image_input_width = 640
    image_input_height = 480

    with gr.Row():
        with gr.Column():
            input_img = gr.Image(sources=["webcam", "upload"], type="pil", width=image_input_width, height=image_input_height, webcam_constraints={"video": {"width": image_input_width, "height": image_input_height}})
        with gr.Column():
            output_img = gr.Image(streaming=True, width=image_input_width, height=image_input_height)            
        with gr.Column():
            table_summary = gr.DataFrame(
                headers=["Box", "Class", "Body part", "Coordinates", "Confidence %"]
            )
            
        dep = input_img.stream(pose_det_cv, input_img, [output_img, table_summary],
                                time_limit=30, stream_every=0.1, concurrency_limit=30)

if __name__ == "__main__":
    demo.queue().launch(debug=False, server_name="0.0.0.0", server_port=7860, share=False)