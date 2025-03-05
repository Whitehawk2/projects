import os
import logging
import socket
from flask import Flask, request
from dotenv import load_dotenv

# Configure logging
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logger = logging.getLogger(__name__)

app = Flask(__name__)

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def hello_handler(path):
    # Get the client's IP address. Handle X-Forwarded-For.
    ip = request.remote_addr
    x_forwarded_for = request.headers.get('X-Forwarded-For')
    
    if x_forwarded_for:
        ip = x_forwarded_for  # Use the first IP from X-Forwarded-For
        # In a production environment, you might want more sophisticated parsing
        # of X-Forwarded-For, as it can contain multiple comma-separated IPs.
    else:
        # If X-Forwarded-For is not present, parse remote_addr to remove the port
        try:
            host, _ = ip.split(':')
            ip = host
        except ValueError:
            # remote_addr doesn't contain a port
            pass
    
    logger.info(f"Request: {request.method} {request.path} from {ip}")
    
    # Write the response to the browser
    return f"Hello {ip}\n"

if __name__ == "__main__":
    # Load environment variables
    if not load_dotenv():
        logger.fatal("Error loading .env file")
        exit(1)
    
    port_str = os.environ.get('LISTEN_PORT')
    if not port_str:
        logger.fatal("LISTEN_PORT environment variable not set.")
        exit(1)
    
    try:
        port = int(port_str)
    except ValueError:
        logger.fatal(f"Invalid port number: {port_str}")
        exit(1)
    
    if port < 1 or port > 65535:
        logger.fatal(f"Invalid port number: {port}. Port must be between 1 and 65535.")
        exit(1)
    
    logger.info(f"Listening on port {port}")
    
    # Start the HTTP server
    app.run(host='0.0.0.0', port=port) 