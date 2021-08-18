import { BaseDataAPI } from "src/app/interface/base.interface";

export interface MyKListDataApi extends BaseDataAPI  {
  data: OneMyKDataBriefResp[]
  status: any[]
}
export interface OneKDataBriefBaseResp {
  UAvatarUrl: string
  NikeName: string
  KCoverImgUrl: string
  KTitle: string
  KBrief: string
}
export interface OneMyKDataBriefResp extends OneKDataBriefBaseResp {

}
