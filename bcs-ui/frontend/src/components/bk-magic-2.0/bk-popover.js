export default {
    name: 'bk-popover',
    computed: {
        instance () {
            return this.$refs.bcsPopover?.instance
        }
    },
    render (h) {
        const slots = Object.keys(this.$slots)
            .reduce((arr, key) => arr.concat(this.$slots[key]), [])
            .map(vnode => {
                vnode.context = this._self
                return vnode
            })
        return h('bcs-popover', {
            ref: 'bcsPopover',
            on: this.$listeners, // 透传事件
            props: Object.assign({}, this.$props, this.$attrs), // 透传props
            scopedSlots: this.$scopedSlots, // 透传scopedSlots
            attrs: this.$attrs // 透传属性，非props
        }, slots)
    }
}
