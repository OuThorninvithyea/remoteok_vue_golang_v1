import { ref, onMounted } from 'vue'
import axios from 'axios'

export function usePostJobPromotion() {
    const promotionData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/promotion')
            promotionData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { promotionData, loading, error }
}
