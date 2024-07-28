<template>
    <div class="text-center">
        <v-dialog v-model="dialog" width="500px">
            <template v-slot:activator="{ on, attrs }">
                <v-btn class="grey darken-1 white--text" v-bind="attrs" v-on="on">
                    <span text color="blue" dense>PREVIEW</span>
                </v-btn>
            </template>
            <v-card>
                <v-card-title class="text-h6 grey lighten-2 text-center">
                    <span><v-btn text black><v-icon class="mr-2 rounded-circle" @click="dialog = false">mdi-close</v-icon>PREVIEW</v-btn></span>
                </v-card-title>

 
                <v-card-text>
                    <v-layout>
                        <v-flex md-12 text-center class="text-h5 mt-3 black--text text--lighten-4 font-weight-bold "><span class="grey--text text--darken-1">MED</span>-APP</v-flex>
                    </v-layout>
                    <v-layout class="ma-3 blue lighten-1 rounded-lg">
                        <v-flex class="ma-3 white--text text-subtitle-1 font-weight-bold">Medicine name</v-flex>
                        <v-flex class="ma-3 white--text text-subtitle-1 font-weight-bold">Qty</v-flex>
                        <v-flex class="ma-3 white--text text-subtitle-1 font-weight-bold">Amount</v-flex>
                    </v-layout>
                    <v-layout v-for="val,idx in Previewarr" :key="idx" class="pl-2 pb-3">
                        <v-flex lg4 text-left class="pl-5">{{ val.medicine_name }}</v-flex>
                        <v-flex lg4 text-center>{{ val.quantity }}</v-flex>
                        <v-flex lg4 text-center>{{ val.amount }}</v-flex>
                    </v-layout>
                    <v-layout class="mb-3">
                        <v-flex lg4></v-flex>
                        <v-flex lg4 text-right class="green--text pr-5 subtitle-1"><b>Total</b></v-flex>
                        <v-flex lg4 text-center class="black--text font-weight-medium ">{{ Total }}</v-flex>
                    </v-layout>
                    <v-layout class="mb-3">
                        <v-flex lg4></v-flex>
                        <v-flex lg4 text-right class="red--text pr-5 subtitle-1"><b>GST</b></v-flex>
                        <v-flex lg4 text-center class="black--text font-weight-medium ">{{ GST }}</v-flex>
                    </v-layout>

                    <v-layout class="mb-3">
                        <v-flex lg4></v-flex>
                        <v-flex lg4 text-right class="black--text font-weight-black pr-5 subtitle-1"><b>Net Price</b></v-flex>
                        <v-flex lg4 text-center class="black--text font-weight-medium ">{{ NetPrice }}</v-flex>
                    </v-layout>

                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn dark color="blue" class="mb-3" @click="printto()">PRINT</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>
<script>
export default {
    name: '',
    data() {
        return {
            arr: ['paracetomol', 5, 250],
            billMaster: this.$store.state.BillMaster,
            billDetails: this.$store.state.BillDetails,
            dialog: false,

        }
    },
    props:{
        Previewarr:[],
        Total:Number, 
        GST:Number,
        NetPrice:Number, 
    },
    methods:{
        printto(){ 
          this.dialog = false;

        }
    }

}
</script>