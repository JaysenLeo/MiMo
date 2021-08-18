export interface BaseOnlyDataAPI {
  data: any
}
export interface BaseOnlyStatusAPI {
  status: any
}
export interface BaseDataAPI extends BaseOnlyDataAPI, BaseOnlyStatusAPI  {
  data: any
  status: any[]
}
