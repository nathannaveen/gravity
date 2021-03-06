/**
 * Copyright 2021 Gravitational Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { Store } from 'nuclear-js';
import { Record } from 'immutable';
import * as actionTypes from './actionTypes';

const NavStoreRec = Record({
  settings: [],
  topNav: [],
  userRole: []
});

export default Store({
  getInitialState() {
    return new NavStoreRec();
  },

  initialize() {
    this.on(actionTypes.NAV_ADD_TOP_ITEM, addTopItem);
    this.on(actionTypes.NAV_ADD_SETTING_ITEM, addSettingItem);
    this.on(actionTypes.NAV_ADD_USERROLE_ITEM, addUserRoleItem);
  }
})

function addTopItem(state, item) {
  const items = [...state.topNav, item];
  return state.set('topNav', items);
}

function addSettingItem(state, item) {
  const items = [...state.settings, item];
  return state.set('settings', items);
}


function addUserRoleItem(state, item) {
  const items = [...state.userRole, item];
  return state.set('userRole', items);
}