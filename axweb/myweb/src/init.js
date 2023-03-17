
console.log("init......")
import mitt from 'mitt'
window.emitter = mitt()

window.tmutils = new class{
    constructor(){
    }
}

class AxWs{
    constructor(host, port){
        this.ws = null
        this.isconnect = false
        this.hostname = host
        this.port = port
        this.wsurl = `wss://${this.hostname}:${this.port}/ws`    
        this.wsopen()  
    }

    wsopen(){
        this.ws = new WebSocket(this.wsurl);        
        
        this.ws.onopen = (ev) => {
            console.log("ws onopen")    
            this.isconnect = true
                    
            this.wssend("subscribe", [ {"channel": "tsmcmqtt1234567890"} ] )
        }
        
        this.ws.onerror = (ev) => {
            console.log(`ws err ${ev}`)
            this.isconnect = false
            this.ws.close()
        }
        
        this.ws.onclose = (ev) => {
            console.log(`ws close`)
            this.isconnect = false            
        }

        this.ws.onmessage = (ev, data) => {            
            let jobj = JSON.parse(ev.data)    
            //console.log("mqtt-->", jobj )
            window.emitter.emit('mqtt', jobj )
        }
    }

    wssend( action, channels ){
        let jobj = { "action": action, "multichannel": channels }        
        this.ws.send( JSON.stringify( jobj) )    
    }
}

window.tmws = new AxWs(location.hostname, 3001)


