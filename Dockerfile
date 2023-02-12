# Use the arm32v7/alpine image as the base image
FROM arm32v7/alpine

# Copy the file to the root directory of the image
COPY arrow_food_api /

ENTRYPOINT ["/arrow_food_api"]