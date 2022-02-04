export let rules = {
    required: value => !!value || 'Required',
    pan: v => new RegExp('[0-9]{8,20}').test(v) || '8-20 digits',
    cardValid: v => new RegExp('[0-9]{2}/[0-9]{2}').test(v) || 'Date in format mm/yy',
  }