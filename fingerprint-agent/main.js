const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const WebSocket = require('ws');
const fetch = require('node-fetch');
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
    // Function to fetch registered users from database
    async function fetchRegisteredUsers() {
      try {
        const response = await fetch('http://localhost:8081/api/v1/users');
        if (response.ok) {
          const data = await response.json();
          return data.data || [];
        }
      } catch (error) {
        console.log('Could not fetch users from database:', error.message);
        return [];
      }
      
      return [];
    }

    // Cache for registered users
    let registeredUsers = [];
    
    // Load registered users on connection
    fetchRegisteredUsers().then(users => {
      registeredUsers = users.filter(user => user.fingerprint_id).map(user => ({
        name: user.name,
        fingerprint_id: user.fingerprint_id
      }));
      console.log(`Loaded ${registeredUsers.length} registered users with fingerprint from database`);
      mainWindow.webContents.send('update-status', { 
        text: `${registeredUsers.length} users loaded from database`, 
        color: 'blue' 
      });
    });

    // Simulate fingerprint scanning with quality variations
    function generateScanResult(user) {
      // Simulate quality variations (85%, 78%, 72%)
      const qualities = [85, 78, 72];
      const quality = qualities[Math.floor(Math.random() * qualities.length)];
      
      // Generate simple template with fingerprint_id
      const template = `TEMPLATE_${user.fingerprint_id}_Q${quality}_${Date.now()}`;
      
      return {
        template: template,
        quality: quality,
        fingerprint_id: user.fingerprint_id
      };
    }

    const sdk = {
      initialize: () => {
        console.log("SDK: Initializing scanner...");
        // TODO: Add your Secugen SDK initialization logic here.
        // Throw an error if initialization fails.
        console.log("SDK: Scanner initialized successfully.");
        mainWindow.webContents.send('update-status', { text: 'Scanner Ready - Place finger on sensor', color: 'green' });
        return true;
      },
      startScan: (callback) => {
        console.log("SDK: Starting scan...");
        mainWindow.webContents.send('update-status', { text: 'Scanning... Place finger on sensor', color: 'orange' });
        
        // Simulate realistic fingerprint scanning with quality variations
        setTimeout(() => {
          if (registeredUsers.length === 0) {
            console.log('No registered users found');
            callback(new Error('No registered users available for scanning'));
            return;
          }
          
          // Simulate random user fingerprint detection from registered users
          const randomUser = registeredUsers[Math.floor(Math.random() * registeredUsers.length)];
          const scanResult = generateScanResult(randomUser);
          
          console.log(`SDK: Fingerprint captured for user: ${randomUser.name}`);
          console.log(`SDK: Fingerprint ID: ${randomUser.fingerprint_id}`);
          console.log(`SDK: Template: ${scanResult.template}`);
          console.log(`SDK: Quality: ${scanResult.quality}%`);
          
          mainWindow.webContents.send('update-status', { 
            text: `Detected: ${randomUser.name} (Q:${scanResult.quality}%)`, 
            color: scanResult.quality >= 80 ? 'green' : scanResult.quality >= 75 ? 'orange' : 'yellow'
          });
          
          // Return the scan result for matching
          callback(null, {
            template: scanResult.template,
            fingerprint_id: randomUser.fingerprint_id,
            user_name: randomUser.name,
            quality: scanResult.quality
          });
        }, 2000 + Math.random() * 1000); // 2-3 seconds
      },
      stopScan: () => {
        console.log("SDK: Stopping scan...");
        mainWindow.webContents.send('update-status', { text: 'Scanner Ready', color: 'green' });
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
      
      // Handle attendance scanning (existing users)
      if (message.command === 'start_scan') {
        sdk.startScan((err, scanResult) => {
          if (err) {
            console.error("Scan error:", err);
            mainWindow.webContents.send('update-status', { text: 'Scan Error', color: 'red' });
            return;
          }
          if (ws.readyState === ws.OPEN) {
            ws.send(JSON.stringify({ 
              type: 'fingerprint_scanned', 
              template: scanResult.template,
              fingerprint_id: scanResult.fingerprint_id,
              user_name: scanResult.user_name,
              quality: scanResult.quality,
              timestamp: new Date().toISOString()
            }));
            console.log(`Sent scan result to frontend:`);
            console.log(`- User: ${scanResult.user_name}`);
            console.log(`- Template: ${scanResult.template}`);
            console.log(`- Quality: ${scanResult.quality}%`);
          }
        });
      }
      
      // Handle enrollment scanning (new users)
      else if (message.command === 'start_enrollment_scan') {
        let enrollmentProgress = 0;
        const maxScans = 3;
        
        const performEnrollmentScan = () => {
          enrollmentProgress++;
          console.log(`Enrollment scan ${enrollmentProgress}/${maxScans}`);
          mainWindow.webContents.send('update-status', { 
            text: `Enrollment: Scan ${enrollmentProgress}/${maxScans}`, 
            color: 'orange' 
          });
          
          // Send progress update
          if (ws.readyState === ws.OPEN) {
            ws.send(JSON.stringify({
              type: 'enrollment_progress',
              progress: enrollmentProgress,
              total: maxScans
            }));
          }
          
          // If we've completed all scans, finish enrollment
          if (enrollmentProgress >= maxScans) {
            console.log('Enrollment completed successfully');
            mainWindow.webContents.send('update-status', { 
              text: 'Enrollment Complete!', 
              color: 'green' 
            });
            
            if (ws.readyState === ws.OPEN) {
              ws.send(JSON.stringify({
                type: 'enrollment_scan_complete',
                message: 'Fingerprint enrollment completed successfully',
                scans_completed: maxScans
              }));
            }
          } else {
            // Schedule next scan after 2 seconds
            setTimeout(performEnrollmentScan, 2000);
          }
        };
        
        // Start first enrollment scan
        console.log('Starting fingerprint enrollment process...');
        mainWindow.webContents.send('update-status', { 
          text: 'Starting Enrollment...', 
          color: 'orange' 
        });
        setTimeout(performEnrollmentScan, 1000);
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