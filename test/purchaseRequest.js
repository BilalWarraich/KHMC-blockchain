const mongoose = require('mongoose');

const PurchaseRequestSchema = new mongoose.Schema({
  requestNo: {
    type: String,
  },
  generatedBy: {
    type: String,
    required: [true, 'Please add generated By'],
  },
  status: {
    type: String,
    required: true,
  },
  committeeStatus: {
    type: String,
    required: true,
  },
  availability:{
    type:Boolean,
    default:true
  },
  reason: {
    type: String,
    required: true,
  },
  vendorId: {
    type: mongoose.Schema.ObjectId,
    ref: 'Vendor',
  },
  rr:{
    type: mongoose.Schema.ObjectId,
    ref: 'ReplenishmentRequest' 
},
  item: {
    itemId: {
      type: mongoose.Schema.ObjectId,
      ref: 'Item',
    },
    currQty: {
      type: Number
    },
    reqQty: {
      type: Number
    },
    comments: {
      type: String
    },
    name: {
      type: String
    },
    description: {
      type: String
    },
    itemCode: {
      type: String
    },
    status: {
      type: String
    },
    secondStatus: {
      type: String
    },
  },
  requesterName: {
    type: String,
  },
  rejectionReason:{
    type:String
  },
  department: {
    type: String,
  },
  commentNotes:{
    type:String
  },
  orderType: {
    type: String,
  },
  generated: {
    type: String,
  },
  approvedBy:{
    type: mongoose.Schema.ObjectId,
    ref: 'staff',
  },
  createdAt: {
    type: Date,
    default: Date.now,
  },
  updatedAt: {
    type: Date,
    default: Date.now,
  },
});

module.exports = mongoose.model('PurchaseRequest', PurchaseRequestSchema);
