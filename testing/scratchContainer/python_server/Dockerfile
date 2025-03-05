FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY main.py .
COPY .env .

# Expose the port (this is metadata, but good practice)
EXPOSE 8080

# Run the application
CMD ["python", "main.py"] 