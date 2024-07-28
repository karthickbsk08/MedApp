<template>
    <div id="managerIV2">
        <v-container>
            
               <h3 class="blue--text darken-3--text font-weight-black text-center mt-5 display-1 text-uppercase">Current Inventry Value</h3>
               <h1 class="grey--text darken-6--text font-weight-bold text-center mt-5 display-3">{{ MgInventoryValue }}</h1>
            
        </v-container>
    </div>
    </template>
    
    <script>
    export default{
        name:'managerIV2',
        data(){
            return{
               
            }
        },
        props:{
            MgInventoryValue:Number
        },
        methods:{
            // getInventoryvalue(){
            //     for(let i=0;i<this.Stockarr.length;i++){

            //        for(let j=0;j<this.MedicineMaster.length;j++){
            //         if(this.Stockarr[i].MedicineName==this.MedicineMaster[j].MedicineName)
            //             this.Todaysales+=(this.Stockarr[i].quantity)*(this.Stockarr[i].Unitprice);
            //        }
            //     }
            // }
        },
       
    }
    
    
    </script>