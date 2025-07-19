const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const { WebSocketServer } = require('ws');

let mainWindow;

function createWindow() {
  mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'),
      contextIsolation: true,
      nodeIntegration: false,
    },
  });

  mainWindow.loadFile('index.html');

  // Open the DevTools.
  // mainWindow.webContents.openDevTools();
}

app.whenReady().then(() => {
  createWindow();

  // Setup WebSocket Server
  const wss = new WebSocketServer({ port: 8088 });
  console.log('WebSocket server started on port 8088');
  mainWindow.webContents.send('update-status', { text: 'Waiting for connection...', color: 'orange' });


  wss.on('connection', function connection(ws) {
    console.log('Frontend connected');
    mainWindow.webContents.send('update-status', { text: 'Connected', color: 'green' });

    // Placeholder for SDK integration
    const sdk = {
      initialize: () => {
        console.log("SDK: Initializing scanner...");
        // TODO: Add your Secugen SDK initialization logic here.
        // Throw an error if initialization fails.
        console.log("SDK: Scanner initialized successfully.");
        return true;
      },
      startScan: (callback) => {
        console.log("SDK: Starting scan...");
        // TODO: Implement the logic to start a scan.
        // This should be an asynchronous operation.
        // When a fingerprint is captured, call the callback with the template.
        // For demonstration, we'll simulate a scan after 2 seconds.
        setTimeout(() => {
          const mockTemplate = "SDK_GENERATED_TEMPLATE_" + Date.now();
          console.log("SDK: Fingerprint captured.");
          callback(null, mockTemplate);
        }, 2000);
      },
      stopScan: () => {
        console.log("SDK: Stopping scan...");
        // TODO: Add logic to stop or release the scanner.
      }
    };

    try {
      sdk.initialize();
      mainWindow.webContents.send('update-status', { text: 'Scanner Ready', color: 'green' });
    } catch (error) {
      mainWindow.webContents.send('update-status', { text: 'Scanner Error', color: 'red' });
      console.error("SDK Initialization failed:", error);
      return;
    }

    ws.on('message', function message(data) {
      const message = JSON.parse(data);
      console.log('Received command: %s', message.command);
      
      if (message.command === 'start_scan') {
        sdk.startScan((err, template) => {
          if (err) {
            console.error("Scan error:", err);
            return;
          }
          if (ws.readyState === ws.OPEN) {
            ws.send(JSON.stringify({ type: 'fingerprint_scanned', template: template }));
          }
        });
      }
    });

    ws.on('close', () => {
      console.log('Frontend disconnected');
      sdk.stopScan();
      mainWindow.webContents.send('update-status', { text: 'Disconnected', color: 'red' });
    });
  });

  app.on('activate', function () {
    if (BrowserWindow.getAllWindows().length === 0) createWindow();
  });
});

app.on('window-all-closed', function () {
  if (process.platform !== 'darwin') app.quit();
});