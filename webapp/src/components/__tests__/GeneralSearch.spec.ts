import { describe, it, expect } from 'vitest'

import { mount } from '@vue/test-utils'
import GeneralSearch from '../GeneralSearch.vue'

import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'

const vuetify = createVuetify({
  components,
})

describe('GeneralSearch', () => {
  it('renders properly', () => {
    const str = 'Test search';
    const wrapper = mount(GeneralSearch, { 
      props: { modelValue: str },
      global: {
        plugins: [vuetify],
      }
    })

    const usernameField = wrapper.find("input");
    expect(usernameField.element.value).toBe(str);
  })
})
