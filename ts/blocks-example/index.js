// Generated by LiveScript 1.6.0
(function(){
  var electron, path, app, BrowserWindow, mainWindow, createWindow;
  electron = require('electron');
  path = require('path');
  app = electron.app;
  BrowserWindow = electron.BrowserWindow;
  mainWindow = {};
  createWindow = function(){
    var mainWindow;
    mainWindow = new BrowserWindow({
      width: 1200,
      height: 700
    });
    mainWindow.loadFile('index.html');
    mainWindow.webContents.openDevTools();
    return mainWindow.on('closed', function(){
      var mainWindow;
      return mainWindow = null;
    });
  };
  app.on('ready', createWindow);
  app.on('window-all-closed', function(){
    if (process.platform !== 'darwin') {
      return app.quit();
    }
  });
  app.on('activate', function(){
    if (mainWindow === null) {
      return createWindow();
    }
  });
}).call(this);
