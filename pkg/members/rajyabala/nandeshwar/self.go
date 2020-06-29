package nandeshwar

import (
  "bhuyan.io/family-tree/pkg/types"
  "bhuyan.io/family-tree/pkg/members/rajyabala/nandeshwar/trinanda"
)

var (
  // Nandeshwar is a member of the family.
  Nandeshwar = types.Member{
    Name: "Nandeswar-Chandana",
    Children: []*types.Member{&trinanda.Trinanda, },
    Gender: types.Male,
  }
)
