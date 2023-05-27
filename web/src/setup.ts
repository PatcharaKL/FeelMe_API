//@ts-ignore
import nodeFetch from 'node-fetch'
import { server } from './mocks/server'
import { cleanup } from '@testing-library/react'

global.Request = nodeFetch.Request
global.fetch = nodeFetch

beforeAll(() => server.listen({ onUnhandledRequest: 'error' }))
afterEach(() => {
    server.resetHandlers();
    cleanup();
})
afterAll(() => server.close())