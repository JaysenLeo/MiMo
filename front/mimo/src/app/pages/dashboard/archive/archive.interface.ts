import { BaseOnlyStatusAPI } from "src/app/interface/base.interface";



export interface ArchiveMeta {
    type?:           '1' | '2' | '3'     // 1 markdown str 2 富文本 str 10 md id
    create_datetime: string   // 创建时间
    thumbs_up_cnt:   number     // 点赞数
    read_cnt:        number
    title:           string             // 标题
}
export interface ArchiveRespData {
    meta:    ArchiveMeta,
    content: string
}
export interface ArchiveRespAPI extends BaseOnlyStatusAPI {
    data: ArchiveRespData
}
