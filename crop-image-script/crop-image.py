from PIL import Image
import sys
import os

def crop_image(image_path, coordinates, output_path):
    # Open the image
    image = Image.open(image_path)
    width, height = image.size

    left    = (params[0] - params[2] / 2) * width
    upper   = (params[1] - params[3] / 2) * height
    right   = (params[0] + params[2] / 2) * width
    lower   = (params[1] + params[3] / 2) * height

    coordinates = (left, upper, right, lower)

    # Crop the image
    cropped_image = image.crop(coordinates)

    # Save the cropped image
    cropped_image.save(output_path)

# Getting image path from user
image_path = input("Enter the path to the directory: ")

# Check if the image exists
if not os.path.exists(image_path):
    print(f"Directory '{image_path}' does not exist.")
    sys.exit()

coordinates_str = input("Enter coordinates (x, y, width, height): ")

params = tuple(map(float, coordinates_str.split(',')))

# Getting output path from user
output_path = input("Enter output path: ")

crop_image(image_path, params, output_path)