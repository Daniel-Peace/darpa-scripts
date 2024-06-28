from ultralytics import YOLO
from PIL import Image
import json

if __name__ == "__main__":
    model = YOLO("/home/pacodataco/runs/detect/yolov8n_custom/weights/best.pt")
    image_path = input("Enter a path to an image:\n")
    image = Image.open(image_path)
    results = model.predict(source=image)
    for result in results:
        combined_data = [{"class": int(cls), "confidence": float(conf)} for cls, conf in zip(result.boxes.cls.cpu().numpy(), result.boxes.conf.cpu().numpy())]
        json.dumps(combined_data)
        print(combined_data)