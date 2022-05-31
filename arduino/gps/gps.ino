#include <WiFi.h>
#include "BluetoothSerial.h"
#include "ELMduino.h"
#include <HTTPClient.h>
#include <TinyGPS++.h>
#include <SoftwareSerial.h>

SoftwareSerial serial1(23, 22); // RX, TX

// The TinyGPS++ object
TinyGPSPlus gps;
#define ELM_PORT   SerialBT
BluetoothSerial SerialBT;
ELM327 myELM327;


// setup
void setup() {
  Serial.begin(115200);
    Serial.println("Esperando por Dados do Modulo...");
  serial1.begin(9600);
  //SerialBT.setPin("1234");
}



// loop
void loop() {
  Serial.print("Entrei  ");
  static unsigned long delayGPS = millis();
  serial1.listen();
   bool lido = false;
 
      while (serial1.available()) {
          char cIn = serial1.read(); 
          Serial.print("\n");
          lido = gps.encode(cIn); 
          if(lido){
            break;
          }
      }

      if (lido) { 
         float flat, flon;
         unsigned long age;
    
        Serial.print("Latitude= "); 
      Serial.print(gps.location.lat(), 6);
      Serial.print(" Longitude= "); 
      Serial.println(gps.location.lng(), 6);



     uint32_t rpm = 0, sped = 0;
  const char* ssid = "AndroidAP"; // Wifi SSID
  const char* password = "ttli5630"; // Wifi Password
  const unsigned int writeInterval = 25000; // write interval (in ms)
  const char* serverName = "http://18.130.231.194:8080/api/auth/trucks/truckState";
  ELM_PORT.begin("ArduHUD", true);


  if (!ELM_PORT.connect("OBDII") || !myELM327.begin(ELM_PORT, true, 2000))
  {
    Serial.println("Couldn't connect to OBD scanner - Phase 1");
    ELM_PORT.end();




    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
    }
    Serial.println("->IP address: ");
    Serial.println(WiFi.localIP());

    HTTPClient http;
    // Your Domain name with URL path or IP address with path
    http.begin(serverName);
    http.addHeader( "Content-Type" , "application/json");
    http.addHeader( "Access-Control-Allow-Origin" , "*");
    // Data to send with HTTP POST
    String httpRequestData = "{ \"truck_id\":\"" + String("2") + "\",\"latitude\":\"" + String(gps.location.lat(), 6) + "\",\"longitude\":\"" + String(gps.location.lng(),6) + "\",\"rpm\":\"" + String(0) + "\",\"speed\":\"" + String(0) + "\"}";
    Serial.println(httpRequestData);
    int httpResponseCode = http.POST(httpRequestData);
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);

    // Free resources
    http.end();
    WiFi.disconnect();
  }










  if (myELM327.status == ELM_SUCCESS)
  {
    Serial.println("Connected to ELM327");

    float tempRPM = myELM327.rpm();
    float tempSPEED = myELM327.kph();
    rpm = (uint32_t)tempRPM;
    sped = (uint32_t)tempSPEED;

    Serial.print("RPM: "); Serial.println(rpm);
    Serial.println(ssid);
    ELM_PORT.end();


    WiFi.begin(ssid, password);
    while (WiFi.status() != WL_CONNECTED) {
      delay(500);
    }
    Serial.println("->IP address: ");
    Serial.println(WiFi.localIP());

    HTTPClient http;
    // Your Domain name with URL path or IP address with path
    http.begin(serverName);
    http.addHeader( "Content-Type" , "application/json");
    http.addHeader( "Access-Control-Allow-Origin" , "*");
    // Data to send with HTTP POST
    String httpRequestData = "{ \"truck_id\":\"" + String("2") + "\",\"latitude\":\"" + String(gps.location.lat(),6) + "\",\"longitude\":\"" + String(gps.location.lng(),6) + "\",\"rpm\":\"" + String(rpm) + "\",\"speed\":\"" + String(sped) + "\"}";
    Serial.println(httpRequestData);
    int httpResponseCode = http.POST(httpRequestData);
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);

    // Free resources
    http.end();
    WiFi.disconnect();
    delayGPS=0;
  } else {
    myELM327.printError();
    delayGPS=0;

  }



  }

 





  delay(200);
}


// GPS displayInfo
