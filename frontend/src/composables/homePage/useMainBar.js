import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useMainBar() {
    const mainBar = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/main-bar')
            mainBar.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { mainBar, loading, error }
}
