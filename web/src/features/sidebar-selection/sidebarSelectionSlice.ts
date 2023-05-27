import { createSlice } from '@reduxjs/toolkit'
import type { PayloadAction } from '@reduxjs/toolkit'
import type { RootState } from '../../app/store'

// Define a type for the slice state
interface sidebarSelectionState {
    selectedItemID: number,
}

// Define the initial state using that type
const initialState: sidebarSelectionState = {
    selectedItemID: 1,

}

export const sidebarSelectionSlice = createSlice({
    name: 'sidebarSelection',
    // `createSlice` will infer the state type from the `initialState` argument
    initialState,
    reducers: {
        setSelectedItem: (state, action: PayloadAction<number>) => {
            state.selectedItemID = action.payload
        },
    },
})

export const { setSelectedItem } = sidebarSelectionSlice.actions

// Other code such as selectors can use the imported `RootState` type
export const selectSidebarSelectedItemID = (state: RootState) => state.sidebarSelection.selectedItemID

export default sidebarSelectionSlice.reducer