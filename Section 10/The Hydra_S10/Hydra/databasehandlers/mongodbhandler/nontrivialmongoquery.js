db.Personnel.find({
  "security clearance": {
    $gt: 3
  },
  "position": {
      "$in": ["Mechanic", "Biologist"]
    }
}).pretty();
