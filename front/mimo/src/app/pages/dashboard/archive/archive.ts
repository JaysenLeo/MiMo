
import { BehaviorSubject } from "rxjs"
import { ArchiveRespData } from "./archive.interface"


export class ArchiveEntity {

    private _archive: ArchiveRespData
    private _archive$: BehaviorSubject<ArchiveRespData>
    constructor() {
        this._archive = this._initArchive()
        this._archive$ = new BehaviorSubject(this._archive)
    }
    private _initArchive = (): ArchiveRespData => {
        return {
            meta: {
                thumbs_up_cnt: 0,
                create_datetime: '',
                read_cnt: 0,
                title: '',
                type: '1'
            },
            content: ''
        }
    }
    setArchive = (archive: ArchiveRespData) => {
        this._archive$.next(archive)
        this._archive = archive
    }
    getArchive = () => {
        return this._archive
    }
    getArchiveObservable = () => {
        return this._archive$
    }
}
